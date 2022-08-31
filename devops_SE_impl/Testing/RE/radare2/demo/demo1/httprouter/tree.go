// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package httprouter

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// 对比两个数,小的返回

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// 找到最长的公共前缀 
func longestCommonPrefix(a, b string) int {
	i := 0
	max := min(len(a), len(b))
	for i < max && a[i] == b[i] {
		i++
	}
	return i
}

// 查找通配符号
// Search for a wildcard segment and check the name for invalid characters.
// Returns -1 as index, if no wildcard was found.
func findWildcard(path string) (wilcard string, i int, valid bool) {
	// Find start
	for start, c := range []byte(path) {
		// A wildcard starts with ':' (param) or '*' (catch-all)
		if c != ':' && c != '*' {
			continue
		}


		// Find end and check for invalid characters
		valid = true
		for end, c := range []byte(path[start+1:]) {
			switch c {
			case '/':
				return path[start : start+1+end], start, valid
			case ':', '*':
				valid = false
			}
		}
		return path[start:], start, valid
	}
	return "", -1, false
}

// 计算: 和 *的数量
func countParams(path string) uint16 {
	var n uint
	for i := range []byte(path) {
		switch path[i] {
		case ':', '*':
			n++
		}
	}
	return uint16(n)
}

//最多255个值 
type nodeType uint8
//枚举 实现 
const (
	static nodeType = iota // default
	root  // value --> 1 
	param // value --> 2 
	catchAll // value --> 3
)
//定义node结构

type node struct {
	path      string
    //路径命
	indices   string
    //最短的下一个匹配 
	wildChild bool
    //是否拥有带通配符的儿子节点 
	nType     nodeType
    //节点类型 (可选:static,root,param.catch)
	priority  uint32
    //权重 这个路径下面还有多少个路径节点
	children  []*node
    //儿子节点,多叉树
	handle    Handle
    //要匹配执行的handler函数 
}

// Increments priority of the given child and reorders if necessary
//调整顺序和提高权重
func (n *node) incrementChildPrio(pos int) int {
    // 子切片
	cs := n.children
    // 儿子切片权重+1 
	cs[pos].priority++
    // 
	prio := cs[pos].priority
    //新权重值 
	// Adjust position (move to front)
	newPos := pos
    // 
	for ; newPos > 0 && cs[newPos-1].priority < prio; newPos-- {
		
        // Swap node positions
        //交换节点位置,交换权重大的节点排在切片的前面的位置,
        // 排序算法 
		cs[newPos-1], cs[newPos] = cs[newPos], cs[newPos-1]
	}
    
	// Build new index char string
    // 排序过后pos就不等于 newPos了 ,newPos 在前面--了 
	if newPos != pos { 
        // 交换索引顺序,这些位操作 下标记操作不调试是真的难看 
        // 本来 是 hc 现在索引是 ch 
		n.indices = n.indices[:newPos] + // Unchanged prefix, might be empty

			n.indices[pos:pos+1] + // The index char we move

			n.indices[newPos:pos] + n.indices[pos+1:] // Rest without char at 'pos'
	}
    // 返回索引的新位置
	return newPos
}




// addRoute adds a node with the given handle to the path.
// 添加节点
// Not concurrency-safe! 非并发安全
// 方法,方法只能被node*类型的结构调用,是一个比package 更小的可视单位.

//输入 路径 handle(web处理逻辑)

//输出 
func (n *node) addRoute(path string, handle Handle) {
	//变量赋值
    fullPath := path
    //权重加一
	n.priority++

	// Empty tree  加入头节点,头节点用来默认的节点加上变量修改
    // 
	if n.path == "" && n.indices == "" {
        //插入孩子节点

		n.insertChild(path, fullPath, handle)
		//设置节点类型 
        n.nType = root
        //返回 
		return
	}

walk:
	for {
		// Find the longest common prefix.
        //寻找最长公共前缀
		// This also implies that the common prefix contains no ':' or '*'
		// since the existing key can't contain those chars.
		i := longestCommonPrefix(path, n.path)
        
		// Split edge
        //路径长度 大于 最大前缀.  
		if i < len(n.path) {
			child := node{
				path:      n.path[i:],
				wildChild: n.wildChild,
				nType:     static,
                //原节点的索引给新的节点
				indices:   n.indices,
                //建立一个新节点,并且把原节点的孩子给他当儿子
				children:  n.children,
				handle:    n.handle,
				priority:  n.priority - 1,
			}
            //任何新节点去当原节点的儿子.
			n.children = []*node{&child}
			// []byte for proper unicode char conversion, see #65
            // 索引修改为最大前缀后的第一个字母,原节点改索引
			n.indices = string([]byte{n.path[i]})
            //路径修改为最大前缀,原节点改名字 
			n.path = path[:i]
			n.handle = nil
			n.wildChild = false
		}
         
		// Make new node a child of this node
        // 新的节点
		if i < len(path) {
            //路径是去掉最长公共前缀的部分  
			path = path[i:]

			if n.wildChild {
				n = n.children[0]
				n.priority++

				// Check if the wildcard matches
				if len(path) >= len(n.path) && n.path == path[:len(n.path)] &&
					// Adding a child to a catchAll is not possible
					n.nType != catchAll &&
					// Check for longer wildcard, e.g. :name and :names
					(len(n.path) >= len(path) || path[len(n.path)] == '/') {
					continue walk
				} else {
					// Wildcard conflict
					pathSeg := path
					if n.nType != catchAll {
						pathSeg = strings.SplitN(pathSeg, "/", 2)[0]
					}
					prefix := fullPath[:strings.Index(fullPath, pathSeg)] + n.path
					panic("'" + pathSeg +
						"' in new path '" + fullPath +
						"' conflicts with existing wildcard '" + n.path +
						"' in existing prefix '" + prefix +
						"'")
				}
			}
           //第一个非公共前缀字母,拿出来当索引,对应ascii 编号 
			idxc := path[0]
                
			// '/' after param
			if n.nType == param && idxc == '/' && len(n.children) == 1 {
				n = n.children[0]
				n.priority++
				continue walk
			}

			// Check if a child with the next path byte exists
            //检查索引是否已经存在  
			for i, c := range []byte(n.indices) {
				if c == idxc {
                    //对已经有索引的节点,只需做 提高权重的操作 
					i = n.incrementChildPrio(i)
                    // 节点调整为子节点,在儿子里面去插.
					n = n.children[i]
                    //循环继续 
					continue walk
				}
			}

			// Otherwise insert it
            //不存在就插入 
			if idxc != ':' && idxc != '*' {
                
				// []byte for proper unicode char conversion, see #65
				//追加一个索引 
                n.indices += string([]byte{idxc})
                //初始化一个新的节点
				child := &node{}
                // 切片追加元素
				n.children = append(n.children, child)
                // 调整孩子权重,给某个位置的索引调整下权重 
				n.incrementChildPrio(len(n.indices) - 1)
				//当前节点移动到孩子root上 
                n = child
			}

            //插入子节点

			n.insertChild(path, fullPath, handle)
			
            return
		}

		// Otherwise add handle to current node
		if n.handle != nil {
			panic("a handle is already registered for path '" + fullPath + "'")
		}
		n.handle = handle
		return
	}
}



//插入孩子
// 输入 路径 全路径 handle 
func (n *node) insertChild(path, fullPath string, handle Handle) {
	// 循环 
    for {
		// Find prefix until first wildcard
        //查找 * 和: 
		wildcard, i, valid := findWildcard(path)
		if i < 0 { // No wilcard found
			break
		}

		// The wildcard name must not contain ':' and '*'
		if !valid {
			panic("only one wildcard per path segment is allowed, has: '" +
				wildcard + "' in path '" + fullPath + "'")
		}

		// Check if the wildcard has a name
		if len(wildcard) < 2 {
			panic("wildcards must be named with a non-empty name in path '" + fullPath + "'")
		}

		// Check if this node has existing children which would be
		// unreachable if we insert the wildcard here
		if len(n.children) > 0 {
			panic("wildcard segment '" + wildcard +
				"' conflicts with existing children in path '" + fullPath + "'")
		}

		// param
		if wildcard[0] == ':' {
			if i > 0 {
				// Insert prefix before the current wildcard
				n.path = path[:i]
				path = path[i:]
			}

			n.wildChild = true
			child := &node{
				nType: param,
				path:  wildcard,
			}
			n.children = []*node{child}
			n = child
			n.priority++

			// If the path doesn't end with the wildcard, then there
			// will be another non-wildcard subpath starting with '/'
			if len(wildcard) < len(path) {
				path = path[len(wildcard):]
				child := &node{
					priority: 1,
				}
				n.children = []*node{child}
				n = child
				continue
			}

			// Otherwise we're done. Insert the handle in the new leaf
			n.handle = handle
			return
		}

		// catchAll
		if i+len(wildcard) != len(path) {
			panic("catch-all routes are only allowed at the end of the path in path '" + fullPath + "'")
		}

		if len(n.path) > 0 && n.path[len(n.path)-1] == '/' {
			panic("catch-all conflicts with existing handle for the path segment root in path '" + fullPath + "'")
		}

		// Currently fixed width 1 for '/'
		i--
		if path[i] != '/' {
			panic("no / before catch-all in path '" + fullPath + "'")
		}

		n.path = path[:i]

		// First node: catchAll node with empty path
		child := &node{
			wildChild: true,
			nType:     catchAll,
		}
		n.children = []*node{child}
		n.indices = string('/')
		n = child
		n.priority++

		// Second node: node holding the variable
		child = &node{
			path:     path[i:],
			nType:    catchAll,
			handle:   handle,
			priority: 1,
		}
		n.children = []*node{child}

		return
	}
 // for 循环结束 

	// If no wildcard was found, simply insert the path and handle
	n.path = path
	n.handle = handle
}














// Returns the handle registered with the given path (key). The values of
// wildcards are saved to a map.

// If no handle can be found, a TSR (trailing slash redirect) recommendation is
// made if a handle exists with an extra (without the) trailing slash for the
// given path.
//如果没有就返回 /的 handle 

func (n *node) getValue(path string, params func() *Params) (handle Handle, ps *Params, tsr bool) {
// for 循环标签控制退出
walk: // Outer loop for walking the tree
	for {
        // 获取节点路径 赋值给前缀变量 
		prefix := n.path
        // 如果path > 前缀,不是不带符号的匹配
		if len(path) > len(prefix) {
			if path[:len(prefix)] == prefix {
				path = path[len(prefix):]

				// If this node does not have a wildcard (param or catchAll)
				// child, we can just look up the next child node and continue
				// to walk down the tree
				if !n.wildChild {
					idxc := path[0]
					for i, c := range []byte(n.indices) {
						if c == idxc {
							n = n.children[i]
							continue walk
						}
					}

					// Nothing found.
					// We can recommend to redirect to the same URL without a
					// trailing slash if a leaf exists for that path.
					tsr = (path == "/" && n.handle != nil)
					return
				}

				// Handle wildcard child
				n = n.children[0]
				switch n.nType {
				case param:
					// Find param end (either '/' or path end)
					end := 0
					for end < len(path) && path[end] != '/' {
						end++
					}

					// Save param value
					if params != nil {
						if ps == nil {
							ps = params()
						}
						// Expand slice within preallocated capacity
						i := len(*ps)
						*ps = (*ps)[:i+1]
						(*ps)[i] = Param{
							Key:   n.path[1:],
							Value: path[:end],
						}
					}

					// We need to go deeper!
					if end < len(path) {
						if len(n.children) > 0 {
							path = path[end:]
							n = n.children[0]
							continue walk
						}

						// ... but we can't
						tsr = (len(path) == end+1)
						return
					}
                // 返回的handle 设置成节点的handle ,handle不是nil 就返回.
					if handle = n.handle; handle != nil {
						return
					} else if len(n.children) == 1 {
						// No handle found. Check if a handle for this path + a
						// trailing slash exists for TSR recommendation
						n = n.children[0]
						tsr = (n.path == "/" && n.handle != nil) || (n.path == "" && n.indices == "/")
					}

					return

				case catchAll:
					// Save param value
					if params != nil {
						if ps == nil {
							ps = params()
						}
						// Expand slice within preallocated capacity
						i := len(*ps)
						*ps = (*ps)[:i+1]
						(*ps)[i] = Param{
							Key:   n.path[2:],
							Value: path,
						}
					}

					handle = n.handle
					return

				default:
					panic("invalid node type")
				}
			}
		} else if path == prefix {
			// We should have reached the node containing the handle.
			// Check if this node has a handle registered.
			if handle = n.handle; handle != nil {
				return
			}

			// If there is no handle for this route, but this route has a
			// wildcard child, there must be a handle for this path with an
			// additional trailing slash
			if path == "/" && n.wildChild && n.nType != root {
				tsr = true
				return
			}

			if path == "/" && n.nType == static {
				tsr = true
				return
			}

			// No handle found. Check if a handle for this path + a
			// trailing slash exists for trailing slash recommendation
			for i, c := range []byte(n.indices) {
				if c == '/' {
					n = n.children[i]
					tsr = (len(n.path) == 1 && n.handle != nil) ||
						(n.nType == catchAll && n.children[0].handle != nil)
					return
				}
			}
			return
		}

		// Nothing found. We can recommend to redirect to the same URL with an
		// extra trailing slash if a leaf exists for that path
		tsr = (path == "/") ||
			(len(prefix) == len(path)+1 && prefix[len(path)] == '/' &&
				path == prefix[:len(prefix)-1] && n.handle != nil)
		return
	}
}














// Makes a case-insensitive lookup of the given path and tries to find a handler.
// It can optionally also fix trailing slashes.
// It returns the case-corrected path and a bool indicating whether the lookup
// was successful.
func (n *node) findCaseInsensitivePath(path string, fixTrailingSlash bool) (fixedPath string, found bool) {
	const stackBufSize = 128

	// Use a static sized buffer on the stack in the common case.
	// If the path is too long, allocate a buffer on the heap instead.
	buf := make([]byte, 0, stackBufSize)
	if l := len(path) + 1; l > stackBufSize {
		buf = make([]byte, 0, l)
	}

	ciPath := n.findCaseInsensitivePathRec(
		path,
		buf,       // Preallocate enough memory for new path
		[4]byte{}, // Empty rune buffer
		fixTrailingSlash,
	)

	return string(ciPath), ciPath != nil
}










// Shift bytes in array by n bytes left
func shiftNRuneBytes(rb [4]byte, n int) [4]byte {
	switch n {
	case 0:
		return rb
	case 1:
		return [4]byte{rb[1], rb[2], rb[3], 0}
	case 2:
		return [4]byte{rb[2], rb[3]}
	case 3:
		return [4]byte{rb[3]}
	default:
		return [4]byte{}
	}
}




























// Recursive case-insensitive lookup function used by n.findCaseInsensitivePath
func (n *node) findCaseInsensitivePathRec(path string, ciPath []byte, rb [4]byte, fixTrailingSlash bool) []byte {
	npLen := len(n.path)

walk: // Outer loop for walking the tree
	for len(path) >= npLen && (npLen == 0 || strings.EqualFold(path[1:npLen], n.path[1:])) {
		// Add common prefix to result
		oldPath := path
		path = path[npLen:]
		ciPath = append(ciPath, n.path...)

		if len(path) > 0 {
			// If this node does not have a wildcard (param or catchAll) child,
			// we can just look up the next child node and continue to walk down
			// the tree
			if !n.wildChild {
				// Skip rune bytes already processed
				rb = shiftNRuneBytes(rb, npLen)

				if rb[0] != 0 {
					// Old rune not finished
					idxc := rb[0]
					for i, c := range []byte(n.indices) {
						if c == idxc {
							// continue with child node
							n = n.children[i]
							npLen = len(n.path)
							continue walk
						}
					}
				} else {
					// Process a new rune
					var rv rune

					// Find rune start.
					// Runes are up to 4 byte long,
					// -4 would definitely be another rune.
					var off int
					for max := min(npLen, 3); off < max; off++ {
						if i := npLen - off; utf8.RuneStart(oldPath[i]) {
							// read rune from cached path
							rv, _ = utf8.DecodeRuneInString(oldPath[i:])
							break
						}
					}

					// Calculate lowercase bytes of current rune
					lo := unicode.ToLower(rv)
					utf8.EncodeRune(rb[:], lo)

					// Skip already processed bytes
					rb = shiftNRuneBytes(rb, off)

					idxc := rb[0]
					for i, c := range []byte(n.indices) {
						// Lowercase matches
						if c == idxc {
							// must use a recursive approach since both the
							// uppercase byte and the lowercase byte might exist
							// as an index
							if out := n.children[i].findCaseInsensitivePathRec(
								path, ciPath, rb, fixTrailingSlash,
							); out != nil {
								return out
							}
							break
						}
					}

					// If we found no match, the same for the uppercase rune,
					// if it differs
					if up := unicode.ToUpper(rv); up != lo {
						utf8.EncodeRune(rb[:], up)
						rb = shiftNRuneBytes(rb, off)

						idxc := rb[0]
						for i, c := range []byte(n.indices) {
							// Uppercase matches
							if c == idxc {
								// Continue with child node
								n = n.children[i]
								npLen = len(n.path)
								continue walk
							}
						}
					}
				}

				// Nothing found. We can recommend to redirect to the same URL
				// without a trailing slash if a leaf exists for that path
				if fixTrailingSlash && path == "/" && n.handle != nil {
					return ciPath
				}
				return nil
			}

			n = n.children[0]
			switch n.nType {
			case param:
				// Find param end (either '/' or path end)
				end := 0
				for end < len(path) && path[end] != '/' {
					end++
				}

				// Add param value to case insensitive path
				ciPath = append(ciPath, path[:end]...)

				// We need to go deeper!
				if end < len(path) {
					if len(n.children) > 0 {
						// Continue with child node
						n = n.children[0]
						npLen = len(n.path)
						path = path[end:]
						continue
					}

					// ... but we can't
					if fixTrailingSlash && len(path) == end+1 {
						return ciPath
					}
					return nil
				}

				if n.handle != nil {
					return ciPath
				} else if fixTrailingSlash && len(n.children) == 1 {
					// No handle found. Check if a handle for this path + a
					// trailing slash exists
					n = n.children[0]
					if n.path == "/" && n.handle != nil {
						return append(ciPath, '/')
					}
				}
				return nil

			case catchAll:
				return append(ciPath, path...)

			default:
				panic("invalid node type")
			}
		} else {
			// We should have reached the node containing the handle.
			// Check if this node has a handle registered.
			if n.handle != nil {
				return ciPath
			}

			// No handle found.
			// Try to fix the path by adding a trailing slash
			if fixTrailingSlash {
				for i, c := range []byte(n.indices) {
					if c == '/' {
						n = n.children[i]
						if (len(n.path) == 1 && n.handle != nil) ||
							(n.nType == catchAll && n.children[0].handle != nil) {
							return append(ciPath, '/')
						}
						return nil
					}
				}
			}
			return nil
		}
	}

	// Nothing found.
	// Try to fix the path by adding / removing a trailing slash
	if fixTrailingSlash {
		if path == "/" {
			return ciPath
		}
		if len(path)+1 == npLen && n.path[len(path)] == '/' &&
			strings.EqualFold(path[1:], n.path[1:len(path)]) && n.handle != nil {
			return append(ciPath, n.path...)
		}
	}
	return nil
}

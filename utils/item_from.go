package utils

import (
	"fmt"
)

type Item struct {
	ID       int
	Name     string
	Icon     string
	Children []int
}

func FindLevel3(father Item, items []Item) string {
	var curHtml string
	// 父节点到子节点的横线程度计算
	childCount := len(father.Children)
	var lineWidth int
	if childCount < 2 {
		lineWidth = 0
	} else {
		lineWidth = int(1 / float32(childCount) * 100)
	}

	// 有子节点，准备子节点外部div
	if childCount > 0 {
		curHtml += fmt.Sprintf(`
		<i class="line1"></i>
		<div class="tree-lv" style="width:100%%">
			<i class="line0" style="width:%d%%"></i>
			<ul>`, lineWidth)
	}

	for _, cid := range father.Children {
		for _, item := range items {
			if item.ID == cid {
				curHtml += fmt.Sprintf(`
				<li class="tree-count%d">
					<i class="line2"></i>
					<img src="%s" alt="%s" title="%s" />
				</li>`, childCount, item.Icon, item.Name, item.Name)
			}
		}
	}

	if childCount > 0 {
		curHtml += `
			</ul>
		</div>`
	}
	return curHtml
}

func FindLevel2(father Item, items []Item) string {

	childCount := len(father.Children)
	var curHtml string
	var lineWidth int
	if childCount < 2 {
		lineWidth = 0
	} else if childCount == 2 {
		lineWidth = 50 //根节点到下一节点横线长度
	} else {
		lineWidth = 66
	}

	if childCount > 0 {
		curHtml += fmt.Sprintf(`
	    <div class="tree-lv" style="width:100%%">
			<i class="line0" style="width:%d%%"></i>
			<ul>`, lineWidth)
	}

	for _, cid := range father.Children {
		for _, item := range items {
			if item.ID == cid {
				curHtml += fmt.Sprintf(`
				<li class="tree-count%d">
					<i class="line2"></i>
					<img src="%s" alt="%s" title="%s" />
					`, childCount, item.Icon, item.Name, item.Name)
				curHtml += FindLevel3(item, items)
			}
		}
		curHtml += `
				</li>`
	}

	if childCount > 0 {
		curHtml += `
			</ul>
		</div>`
	}
	return curHtml
}

func ItemFormTree(item Item, allItems []Item) string {
	var finalHtml string

	// 头部
	finalHtml = `
	<html>
    <head>
        <link rel="stylesheet" href="//lol.qq.com/c/=/v3/css/topfoot.css,/v3/css/comm.css,/v3/css/v2public.css" />
    </head>
    <body>
        <div id="itemFromTree" class="item-from">
            <div class="clearfix item-tree item-depth3">`

	//本身
	finalHtml += fmt.Sprintf(`
		<div class="tree-lv"> 
			<img src="%s" alt="%s" title="%s" /> 
			<i class="line1"></i> 
	    </div>`, item.Icon, item.Name, item.Name)

	finalHtml += FindLevel2(item, allItems)

	tail := `
			</div>
		</div>
	</body>
	</html>`

	finalHtml += tail
	return finalHtml
}

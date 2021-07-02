;(function($) {
    if($==null) {


        console.warn("没有导入 jquery  qwq..")
        return
    }
    // function jsonAdaptor(json) {
    //     if(json==null) {
    //         return {}
    //     }
    //     //解决go 后台序列化的 坑
    //     for(let name in json) {
    //         // console.log(name)
    //         // console.log(json)
    //         var val = json[name];
    //         if(val =="true" ) {
    //             json[name] = true
    //         }else if(val == "false" ) {
    //             json[name] = false
    //         }else if(val == null ) {
    //
    //         }else if(isNaN(val)) {
    //             json[name] = val +''
    //         }/*else {
    //             json[name] = parseFloat(val)
    //
    //         }*/
    //     }
    //     // console.log(json)
    //     return json
    // }


    //防止重复点击
    const RequestOne = {};
    $.extend({
        //封装 ajax 请求
        Request: function(config) {
            config.type = config.type || config.method || config.methods||'GET';
            if(config.url == null||config.url ==='') {
                console.warn('没有输入 URL')
            }
            let key = config.key||'request_one';
           // console.log(RequestOne[key])
            if (RequestOne[key]) return;
            RequestOne[key] = true;
            let callback = config.complete;
            config.complete = function (e) {
                RequestOne[key] = null;
                // config.complete.bind(this)(e)
                if(callback) {
                    callback.bind(this)(e)
                }
            }
            // config.url = config.url||''

            config.data  =  JSON.stringify( (config.data) )
            config.contentType = config.contentType||"application/json;charset=utf-8"
            config.dataType = config.dataType||'json'
            config.xhrFields = config.xhrFields|| {withCredentials: true}
            config.success = config.success|| config.ok
            config.error = config.error||config.fail
            config.complete = config.complete || config.onComplete ||config.then
            return $.ajax(config)
        },
        RequestFormData: function(config) {
            config.type = config.type || config.method || config.methods||'POST';
            config.url = config.url||''
            config.data  =  config.data||config.params||config.param||config.form
            config.processData = false

            config.contentType = false
            config.xhrFields = config.xhrFields|| {withCredentials: true}
            config.success = config.success|| config.ok
            config.error = config.error||config.fail
            config.complete = config.complete || config.onComplete ||config.then

            return $.ajax(config)
        },
        //表单转 JSON
        FormToJson: function(ele) {
            let o = {};
            const el = $(ele)
            let a = el.serializeArray();
            $.each(a, function() {
                const dom  = $(`[name=${this.name}]`)
                let type = dom.data('type')||'string'
                // console.log(dom[0])
                // console.log(type)
                // console.log(this)
                let currentValue = this.value

                // console.log(currentValue)
                if(type=="string" || type=='str') {
                    currentValue = currentValue+''
                }
                else if( type == 'int' || type=="integer") {
                    try{
                        currentValue = parseInt(currentValue)
                    }catch(e) {
                        log.warn("转化失败..")
                        currentValue = null
                    }

                }else if(type == 'boolean' || type == 'bool' ) {
                    if(currentValue==0 || currentValue==null||type=='false') currentValue = false;
                    else currentValue = true;
                }else if(type == 'float' || type=='number') {
                    try{
                        currentValue = parseFloat(currentValue)
                    }catch (e) {
                        log.warn("转化失败.float")
                        currentValue = null
                    }

                }
                if(currentValue=="undefined" || currentValue==undefined) {
                    currentValue = null
                }
                if (o[this.name]) {
                    if (!o[this.name].push) {
                        o[this.name] = [ o[this.name] ];
                    }
                    o[this.name].push(currentValue);
                } else {
                    o[this.name] = currentValue
                }
            });
            return o;
        },
        //JSON 转 URL
        JsonToQuery: function(json) {
            return  Object.keys(json).map(function (key) {
                // body...
                return encodeURIComponent(key) + "=" + encodeURIComponent(json[key]);
            }).join("&");
        },
        _name_val: function(name) {
            const express = `*[name="${name}"]`
            // const el = _$(express)
            return $(express).val()||''
        }


    })



    //娉ㄥ唽 dom 鍑芥暟 , 閫氳繃 name 鏉ヨ幏鍙栬〃鍗曞睘鎬�
    let __findByName = function(name) {
        //鎵剧1涓� name=鏌愭煇鏌愮殑 dom
        return this.find(`*[name=${name}]`).eq(0)
    }
    $.fn.findByName = $.fn.queryByName = $.fn.selectByName = __findByName;




    //-----------


    //鍒嗛〉鎻掍欢
    function _pages(curPage,totalPage) {
        const c = curPage
        const t = totalPage
        if (t<10) {
            let a = [];
            for (let i=1;i<= totalPage;++i) a.push(i);
            return a;
        }
        if (c <= 5) {
            return [1, 2, 3, 4, 5, 6, 7, 8, 9, '...', t]  //绗竴绉嶆儏鍐�
        } else if (c >= t - 4) {
            return [1, '...', t-8, t-7, t-6, t-5, t-4, t-3, t-2, t-1, t] //绗簩绉嶆儏鍐�
        } else {
            return [1, '...', c-3, c-2, c-1, c, c+1, c+2, c+3, '...', t]  //绗笁绉嶆儏鍐�
        }
    }

    function _isNum(n) {
        return !isNaN(parseFloat(n)) && isFinite(n);

    }


    function _renderPagination(curPage,totalPage,callback) {
        let args = arguments;
        if(args.length!=3) {
            console.warn('鍒嗛〉鏈夐棶棰�,curPage,totalPage ')
        }
        let pageDiv = $('.pagination')
        let arr = _pages(curPage,totalPage)
        let str = ""
        for (let item of arr) {

            if (_isNum(item)) {

                if(curPage==item) {
                    str = str+`<li class='active  page-item' href="#${item}" data-page="${item}"><a class="page-link">${item}</a></li>`
                }else
                {
                    str = str+`<li  class="page-item" data-page="${item}"><a class="page-link">${item}</a></li>`
                }
            }else {
                str = str+"<li  class='page-item disabled'> <a class='page-link'>"+item+"</a></li>"
            }
        }
        pageDiv.html(str)
        pageDiv.find('li').on('click',function (e) {
            if(callback) callback.bind(this)(e)
        })
    }
    //鍒嗛〉鍑芥暟
    $.extend({
        renderPage: _renderPagination
    })


})(window.jQuery)

;




//封装 获取表单元素
(function($) {
    //被选中 了的
    function _selected() {
        let res = [];
        this.each(function() {
            if($(this).prop('checked')) {
                res.push($(this)[0])
            }
        })
        return $(res)

    }
    //value 数组
    function _valueArr() {
        let res = []
        $(this).each(function() {
            res.push($(this).val()||'')
        })
        return res
    }

    //属性数组
    function _attrArr(attrName) {
        if(attrName==null) {
            console.warn('没有属性名字，qwq. attrArr')
            return [];
        }
        let res = []
        $(this).each(function() {
            res.push($(this).attr(attrName))
        })
        return res;
    }

    $.fn.extend({
        selected: _selected,
        checked: _selected,
        valueArr: _valueArr,
        valueArray: _valueArr,
        attrArr: _attrArr,
        attrArray: _attrArr,
    })
})(window.jQuery);


(function($) {
    //自动转化类型
    function convertType(json,config) {
        if (json==null) return {}
        if(config==null) return json
        for(var k in config) {
            let type = config[k] ||'str'
            if (type=='str' || type=='string') {
                json[k] = json[k]||''
            }else if (type=='int') {
                try{
                    json[k] = parseInt(json[k])
                }catch (e) {
                    json[k] = null
                    log.warn('类型转化失败')
                }
            }else if(type=='number' || type=='float') {
                try{
                    json[k] = parseFloat(json[k])
                }catch (e) {
                    json[k] = null
                }
            }else if(type=='bool') {
                if(json[k]=='false' || json[k] ==null || json[k] == 0) {
                    json[k] = false
                }else {
                    json[k] = true
                }
            }
            //其他类型

        }
        return json
    }
    $.extend({
        convertType:convertType
    })
})(window.jQuery);



//防抖
(function ($) {
    function debounce(handler, delay) {
        let timer;
        return function() {
            clearTimeout(timer)
            let ctx = this
            let args = arguments
            timer = setTimeout(function() {
                handler.apply(ctx, args)
            },delay)

        }
    }
    function throttle(handler,delay) {
        let pre = 0;
        return function() {
            let cur = new Date();

            if(cur - pre > delay) {
                let ctx = this;
                let args = arguments
                handler.apply(ctx, args)
                pre = cur
            }

        }
    }
    $.extend({
        Debounce:debounce,

        Throttle:throttle

    })

})(window.jQuery);
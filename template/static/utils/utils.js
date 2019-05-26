/**
 * 大雄Js 常用utils
 **/

/*
    check 部分
 */
function checkParamNotNullValidate(param) {
    if(param != null && param != '') {
        return true;
    }
    return false;
}

/*
    request 请求封装
 */
function request(url,data,suc,err,method) {
    if (method == "GET") {
        $.ajax({
            type : "GET",
            url : url,
            success : function(data){
                suc(data);
            },
            error : function (data) {
                err(data);
            }
        });
    }else{
        $.ajax({
            type : "POST",
            url : url,
            data : data,
            success : function(data){
                suc(data);
            },
            error : function (data) {
                err(data);
            }
        });
    }
}

////////////////////基础工具//////////////////////////
/*
    obj->json
 */
function objToJson(obj) {
    return JSON.stringify(obj)
}
/*
    json->obj
 */
function jsonToObj(json) {
    return JSON.parse(json)
}
/*
    判断localStorage 值不为空
 */
function localStorageHas(key) {
    let s = localStorage.getItem(key);
    return checkParamNotNullValidate(s)
}
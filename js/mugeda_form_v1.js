// 通用登陆代码开始
const domain = "https://t.iuu.pub";
const search = (variable) => {
  var arrStr = window.location.search.substring(1).split("&");
  for (var i = 0; i < arrStr.length; i++) {
    var temp = arrStr[i].split("=");
    if (temp[0] === variable){
        return decodeURIComponent(temp[1]);
    }  
  }
  return false;
};
window.setCookie = function (name,value,Days){
    // var Days = 30;
    var exp = new Date();
    exp.setTime(exp.getTime() + Days*24*60*60*1000);
    document.cookie = name + "="+ escape (value) + ";expires=" + exp.toGMTString();
};
window.getCookie = function (name){
    arrStr = document.cookie.split("; ");
    for (var i = 0; i < arrStr.length; i++) {
    var temp = arrStr[i].split("=");
    if (temp[0] === name){
        return decodeURIComponent(temp[1]);
    }    
  }
  return false;
};
var wechat_jwt_token = getCookie("wechat_jwt_token");

wechat_oauth = () => {
    setCookie("is_wechat_jwt_token","sasdas");
    // var reuri = window.location.origin+window.location.pathname;
    var reuri = window.location.href;
    fetch(domain+"/v2/api/oauth/wechat?state="+encodeURIComponent(reuri))
    .then((res)=>res.json())
    .then((res)=>{
        if(res.code == 1){
            window.location.href = res.data;
        }else{
            Mugine.Utils.Toast.info( res.msg+"err_msg:"+res.err, {type:'info'});
        }
    })
    .catch((err)=>{
        Mugine.Utils.Toast.info(err, {type:'info'});
    });
}

init_oauth = () => {
  /* 测试环境
    if(search("wechat_jwt_token")){ // 设置token
        setCookie("wechat_jwt_token",search("wechat_jwt_token"),1/48);
    }
    if(!wechat_jwt_token && !search("wechat_jwt_token")){ // 未登陆
        
        wechat_oauth();
    }else{ //
        if(search("wechat_jwt_token")){ // 设置token
            setCookie("wechat_jwt_token",search("wechat_jwt_token"),1/48);
        }
    }
    */
    // 正式环境使用
    
    if(getCookie("is_wechat_jwt_token")){ // 执行 wechat_oauth 就会存在这个
        if(!wechat_jwt_token && !search("wechat_jwt_token")){ // 未登陆
            wechat_oauth();
            return
        }
        if(!wechat_jwt_token){
            if(search("wechat_jwt_token")){ // 登录成功 设置token
                setCookie("wechat_jwt_token",search("wechat_jwt_token"),1/48);
                setCookie("is_wechat_jwt_token","is_wechat_jwt_token",-1);
                wechat_jwt_token = getCookie("wechat_jwt_token"); // 赋值登录信息 
            }else{

            }
        }
    }else{
        wechat_oauth();
    }
    
};
function getIsWxClient () {
  var ua = navigator.userAgent.toLowerCase();
  if (ua.match(/MicroMessenger/i) == "micromessenger") {
      return true;
  }
  return false;
};

// init_oauth();
// 通用登陆代码结束

var userinfo;// obj
var err_msg; // 错误信息
// 显示头像
HeadShow = () => {
  // headimgsrc = URLParam("headimgurl");
  let scene = mugeda.scene;
  // var aObject = scene.getObjectByName("命名物体的名字");
  let headimg = scene.getObjectByName("headimg");
  headimg.src = "https" + userinfo.head_img.substring(4);
  console.log("头像设置成功！");
  // 事件埋码
  yhSDK.event({
    "event_name": "显示头像",//事件名称
    "category": "HeadShow",//事件对象
    "action": "renderready",//事件操作
    "event_param": JSON.stringify({"nick_name": Base64.decode(userinfo.nick_name), "union_id": userinfo.union_id}) //选填参数
  });
};
// 设置微信昵称
SetNickName = () => {
  let scene = mugeda.scene;
  let wx_nick_name = scene.getObjectByName("微信昵称");
  wx_nick_name.text = Base64.decode(userinfo.nick_name)
  yhSDK.event({
    "event_name": "设置微信昵称",//事件名称
    "category": "SetNickName",//事件对象
    "action": "renderready",//事件操作
    "event_param": JSON.stringify({"nick_name": Base64.decode(userinfo.nick_name), "union_id": userinfo.union_id}) //选填参数
  });
}
// YH SDK
sdk_initApp = (union_id) => {
  //initializing yhSDK with params
  let params = {}
  //必填参数 ----------------------- start
  //应用对应的唯一标识,由Tracking System生成并提供，以下为测试与生产环境的测试app_key，
  // for_test: 是否为测试环境 true为是，false为否，默认生产环境；
  params.for_test = true;

  if(params.for_test){
    params.app_key = "qa1575860207903"; // QA测试app_key
  }else{
    params.app_key = "ts1575860189968"; // 正式的app_key在测试完毕后再由Tracking System负责人提供
  }

  //页面唯一标识 命名规则：以'page_' 开始，后续接驼峰命名规则如"page_userCenter"
  params.page_id = "page_index";
  //页面唯一名称
  params.page_name = "首页";
  //必填参数 -----------------------end

  //选填参数 可以不填 ----------------------- start
  //用户唯一标识（未授权或未登录无法获取用户ID的话，则不用填写该参数
  // params.device_id = "open_id";
  //用户唯一标识（多个应用之间唯一的ID，需关注微信公众号
  if(union_id!=""){
    params.union_id = union_id;
  }
 
  //用户唯一标识（授权或登录后可以获取到用户ID的话，将ID赋值给user_id
  // params.user_id = "186XXXXXXXX";

  //业务渠道 应用个性渠道 如“消费者”、“促销员”, 没有相关业务场景可以去掉
  params.business_channel = "inch";
  // 页面通用采集参数
  // params.pageview_param = JSON.stringify({"promo_id": 25});
  //以POST请求发送
  params.force_post = true;
  //选填参数 ----------------------------------------- end
  yhSDK.init(params)
  SetNickName();
  HeadShow();
}
if(getIsWxClient()){
  // 授权登录
  init_oauth();
  
}else{
  userinfo = {
    app_id:"", // wxbdb9cd64895da3d3
    head_img:"",// http://thirdwx.qlogo.cn/mmopen/vi_32/KxvECY1NtxoPDTeAWibu8MdichtmeEdkT0kckkt6e44ujBIdgibaFPjxyK1IwSQZp0pEaN3OwAlsjJJ1UrfL8kHew/132
    nick_name:"", // 6JSa
    open_id:"", // oLzWXwZk8QRu53nv49S5U9B3VCNo
    union_id:"" // oEo3p1bg1EN5-XhbIdMYQV7W1Ul0
  }
  sdk_initApp("")
}
GetUser = (...callbak) => {
  // 获取用户信息 v2/api/oauth/wechat/userinfo
  fetch(domain+"/v2/api/oauth/wechat/userinfo",{
    headers: {
      'wechat_jwt_token': wechat_jwt_token,
    },
    method: 'GET', // *GET, POST, PUT, DELETE, etc.
    mode: 'cors', // no-cors, cors, *same-origin
    redirect: 'follow', // manual, *follow, error
    referrer: 'no-referrer', // *client, no-referrer
  })
  .then((res)=>res.json())
  .then((res)=>{
      if(res.code == 1){
        //Mugine.Utils.Toast.info("{data:"+res.data.union_id, {type:'info'});
        userinfo = res.data;
        /*
        for(let i = 0; i < callbak.length,i++;){
          callbak[i]()
          console.log(callbak[i])
        }
        */
        // 微信内
        sdk_initApp(res.data.union_id);
        console.log(res.data);
      }else{
        Mugine.Utils.Toast.info( res.msg+"err_msg:"+res.err, {type:'info'});
      }
  })
  .catch((err)=>{
      Mugine.Utils.Toast.info(err, {type:'info'});
  });
}

// 初始化表单值
var pic_name="",pic_url="",text="",is_cook=false;

// 是否喜欢烹饪
IsCook = (em,is) => {
  if(is == "1"){
    is_cook=true
  }
  console.log("is_cook:",is_cook)
  // 事件埋码
  yhSDK.event({
    "event_name": "是否喜欢烹",//事件名称
    "category": "IsCook",//事件对象
    "action": "click",//事件操作
    "event_param": JSON.stringify({"nick_name": Base64.decode(userinfo.nick_name), "union_id": userinfo.union_id}) //选填参数
  });
}

// 获取pic_name
GetPicName = (em) =>{
  if(em.x == 51){
    if(pic_name == ""){
      pic_name = em.name
    }else{
      pic_name = pic_name + "+" +em.name
    }
  }
}

// 提交表单
PostForm = (em) => {
  if(!getIsWxClient()){
    return
  }
  let scene = mugeda.scene;
  pic_name = scene.getObjectByName("食物文本").text+"+"+scene.getObjectByName("配料文本").text;
  // pic_url = scene.getObjectByName("容器1").src;
  pic_url = em.src
  console.log(pic_url)
  text = scene.getObjectByName("故事").text;

  let data = {
    pic_name:pic_name,
    pic_url:pic_url.substring(22),
    text:text,
    is_cook:is_cook
  }
  let url = domain+"/v2/mugeda/form/v1"
  const PostFormEvent = (err_msg) => {
    yhSDK.event({
      "event_name": "提交表单",//事件名称
      "category": "PostForm",//事件对象
      "action": "POST",//事件操作
      "event_param": JSON.stringify({"nick_name": Base64.decode(userinfo.nick_name), "union_id": userinfo.union_id,"err_msg":err_msg,"post_url":url,"post_body":data}) //选填参数
    });
    console.log(err_msg,data)
  }

  if(pic_url=="" ||  text==""){
    err_msg = "名称/图片/文字 未填写或上传"
    Mugine.Utils.Toast.info(err_msg, {type:'info'});
    PostFormEvent(err_msg)
    return
  }
  
  fetch(url,{
    headers: {
      'wechat_jwt_token': wechat_jwt_token,
    },
    method: 'POST', // *GET, POST, PUT, DELETE, etc.
    mode: 'cors', // no-cors, cors, *same-origin
    redirect: 'follow', // manual, *follow, error
    referrer: 'no-referrer', // *client, no-referrer
    body:JSON.stringify(data)
  })
  .then((res)=>res.json())
  .then((res)=>{
      if(res.code == 1){
        err_msg = res.msg
        PostFormEvent(err_msg)
        // Mugine.Utils.Toast.info( err_msg, {type:'info'});
        console.log(res.data)
      }else{
        err_msg =  res.msg+"err_msg:"+res.err
        PostFormEvent(err_msg)
        Mugine.Utils.Toast.info(err_msg, {type:'info'}); 
      }
  })
  .catch((err)=>{
    err_msg = err
    PostFormEvent(err_msg)
    Mugine.Utils.Toast.info(err_msg, {type:'info'}); 
  });
}

// 获取表单
GetForm = () => {
  fetch(domain+"/v2/mugeda/form/v1",{
    headers: {
      'wechat_jwt_token': wechat_jwt_token,
    },
    method: 'GET', // *GET, POST, PUT, DELETE, etc.
    mode: 'cors', // no-cors, cors, *same-origin
    redirect: 'follow', // manual, *follow, error
    referrer: 'no-referrer', // *client, no-referrer
  })
  .then((res)=>res.json())
  .then((res)=>{
      if(res.code == 1){
        Mugine.Utils.Toast.info(res.msg, {type:'info'});
        console.log(res.data)
      }else{
        Mugine.Utils.Toast.info( res.msg+"err+msg:"+res.err, {type:'info'});
      }
  })
  .catch((err)=>{
      Mugine.Utils.Toast.info(err, {type:'info'});
  });
}
CheckLogin = () => {  
  if(wechat_jwt_token){
    GetUser()
  }else{
    i+1
    setTimeout("CheckLogin()",200);
    console.log("CheckLogin()",i)
  }
}

mugeda.addEventListener("renderready", function(){
  // 当动画准备完成，开始播放前的那一刻引发回调。
  var i = 0

  CheckLogin()
});
// 通用登陆代码开始
const domain = "https://www.inch.online";
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
wechat_jwt_token = getCookie("wechat_jwt_token_fv2");

wechat_oauth = () => {
    // var reuri = window.location.origin+window.location.pathname;
    var reuri = window.location.href;
    fetch(domain+"/v2/mugeda/form/v2/oauth/wechat?state="+encodeURIComponent(reuri))
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
userinfo="";
// 获取用户信息
GetUserInfo = () => {
  var scene = mugeda.scene;
  window.s = scene
  var head_img = scene.getObjectByName("头像")
  var nick_name = scene.getObjectByName("微信昵称")
    if(!wechat_jwt_token){
      wechat_oauth();
      return
    }
    // 获取用户信息 
    fetch(domain+"/v2/mugeda/form/v2/oauth/wechat/userinfo",{
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
        userinfo = res.data;
        head_img.src = "https" + res.data.head_img.substring(4)
        nick_name.text = Base64.decode(res.data.nick_name)
        console.log(res.data);
      }else{
        Mugine.Utils.Toast.info( res.msg+"err_msg:"+res.err, {type:'info'});
      }
    })
    .catch((err)=>{
      Mugine.Utils.Toast.info(err, {type:'info'});
    });
}
init_oauth = () => {
  if(!wechat_jwt_token){
    if(search("code")){ // 登录成功 设置token
      fetch(domain+"/v2/mugeda/form/v2/oauth/wechat/token?code="+search("code"))
      .then((res)=>res.json())
      .then((res)=>{
          if(res.code == 1){
            var token = res.data
            setCookie("wechat_jwt_token_fv2",token,1/1440);
            wechat_jwt_token = token; // 赋值登录信息 
            // GetUserInfo()
          }else{
            wechat_oauth();
          }
      })
      .catch((err)=>{
        Mugine.Utils.Toast.info(err, {type:'info'});
      });
      
    }else{
      wechat_oauth();
    }
  }
};
function getIsWxClient () {
  var ua = navigator.userAgent.toLowerCase();
  if (ua.match(/MicroMessenger/i) == "micromessenger") {
      return true;
  }
  return false;
};
init_oauth();
// 通用登陆代码结束

// 填写信息
type_message = () => {
  var scene = mugeda.scene;
  if(userinfo.phone == ""){
    scene.gotoPage(5);
  }else{
    Mugine.Utils.Toast.info("您的信息已录入,无需重复操作", {type:'info'});
  }
}

// 提交愿望
submit_desire = () => {
  var scene = mugeda.scene;
  var text = scene.getObjectByName("愿望").text
  if(text.length < 1 ){
    Mugine.Utils.Toast.info("未输入愿望", {type:'info'});
    return
  }
  // 提交愿望 
  fetch(domain+"/v2/mugeda/form/v2?text="+text,{
    headers: {
      'wechat_jwt_token': wechat_jwt_token,
    },
    method: 'POST', // *GET, POST, PUT, DELETE, etc.
    mode: 'cors', // no-cors, cors, *same-origin
    redirect: 'follow', // manual, *follow, error
    referrer: 'no-referrer', // *client, no-referrer
  })
  .then((res)=>res.json())
  .then((res)=>{
    if(res.code == 1){
      scene.gotoAndPlay(5, 1);
      console.log(res.data);
    }else{
      Mugine.Utils.Toast.info( res.msg+"err_msg:"+res.err, {type:'info'});
    }
  })
  .catch((err)=>{
    Mugine.Utils.Toast.info(err, {type:'info'});
  });
}
// 提交个人信息
submit_selfinfo = () => {
  var scene = mugeda.scene; 
  var name = scene.getObjectByName("姓名").text
  if(name.length < 1 ){
    Mugine.Utils.Toast.info("未输入姓名", {type:'info'});
    return
  }
  var phone = scene.getObjectByName("电话").text
  if(phone.length < 3 ){
    Mugine.Utils.Toast.info("未输入电话", {type:'info'});
    return
  }
  var address = scene.getObjectByName("地址").text
  if(address.length < 1 ){
    Mugine.Utils.Toast.info("未输入地址", {type:'info'});
    return
  }
  // 提交个人信息 
  fetch(domain+"/v2/mugeda/form/v2/updates?name="+name+"&phone="+phone+"&address="+address,{
    headers: {
      'wechat_jwt_token': wechat_jwt_token,
    },
    method: 'POST', // *GET, POST, PUT, DELETE, etc.
    mode: 'cors', // no-cors, cors, *same-origin
    redirect: 'follow', // manual, *follow, error
    referrer: 'no-referrer', // *client, no-referrer
  })
  .then((res)=>res.json())
  .then((res)=>{
    if(res.code == 1){
      scene.gotoAndPause(1, 5);
      userinfo = res.data
      console.log(res.data);
    }else{
      Mugine.Utils.Toast.info( res.msg+"err_msg:"+res.err, {type:'info'});
    }
  })
  .catch((err)=>{
    Mugine.Utils.Toast.info(err, {type:'info'});
  });
}

// 我的愿望记录
my_desire = () => {
  var scene = mugeda.scene; 
  var my_desire_obj = scene.getObjectByName("我的愿望记录")
  fetch(domain+"/v2/mugeda/form/v2/finds",{
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
      var str = ""
      for(var i=0;i<res.data.length;i++){
        var n = i+1
        str = str+ n+"、" + res.data[i].text+"\r\n"
      }
      my_desire_obj.text = str
      console.log(res.data);
    }else{
      Mugine.Utils.Toast.info( res.msg+"err_msg:"+res.err, {type:'info'});
    }
  })
  .catch((err)=>{
    Mugine.Utils.Toast.info(err, {type:'info'});
  });
}
mugeda.addEventListener("renderready", function(){
  // 当动画准备完成，开始播放前的那一刻引发回调。
  // var vConsole = new VConsole();
  // console.log('Hello world');
  GetUserInfo()
});
// https://cdn.jsdelivr.net/npm/@sweetalert2/theme-borderless/borderless.css
function dynamicLoadCss(url) {
  var head = document.getElementsByTagName('head')[0];
  var link = document.createElement('link');
  link.type = 'text/css';
  link.rel = 'stylesheet';
  link.href = url;
  head.appendChild(link);
}
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
    fetch(domain+"/v3/mugeda/form/v3/oauth/wechat?state="+encodeURIComponent(reuri))
    .then((res)=>res.json())
    .then((res)=>{
        if(res.code == 1){
            window.location.href = res.data;
        }else{
            Mugine.Utils.Toast.info( res.msg+"err_msg:"+res.err, {type:'info'});
        }
    })
    .catch((err)=>{
      Swal.fire({
        icon: 'error',
        text: err,
        confirmButtonText: "好的"
      })
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
    fetch(domain+"/v3/mugeda/form/v3/oauth/wechat/userinfo",{
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
        // Mugine.Utils.Toast.info( res.msg+"err_msg:"+res.err, {type:'info'});
        Swal.fire({
          icon: 'error',
          title: res.msg,
          text: res.err,
          confirmButtonText: "好的"
        })
      }
    })
    .catch((err)=>{
      Swal.fire({
        icon: 'error',
        text: err,
        confirmButtonText: "好的"
      })
    });
}
init_oauth = () => {
  if(!wechat_jwt_token){
    if(search("code")){ // 登录成功 设置token
      fetch(domain+"/v3/mugeda/form/v3/oauth/wechat/token?code="+search("code"))
      .then((res)=>res.json())
      .then((res)=>{
          if(res.code == 1){
            var token = res.data
            setCookie("wechat_jwt_token_fv2",token,1/1440);
            wechat_jwt_token = token; // 赋值登录信息 
          }else{
            wechat_oauth();
          }
      })
      .catch((err)=>{
        Swal.fire({
          icon: 'error',
          text: err,
          confirmButtonText: "好的"
        })
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


mugeda.addEventListener("renderready", function(){
  scene = mugeda.scene
  // 加载CSS
  dynamicLoadCss('https://cdn.jsdelivr.net/npm/@sweetalert2/theme-borderless/borderless.css')
  // 
  /*
  Swal.fire(
    'Good job!',
    'You clicked the button!',
    'success'
  )
  */
});

// 提交用户信息
put_user_info = () => {
  var name = scene.getObjectByName("姓名").text
  if(name.length < 1 ){
    // Mugine.Utils.Toast.info("未输入姓名", {type:'info'});
    Swal.fire({
      icon: 'error',
      text: "未输入姓名",
      confirmButtonText: "好的"
    })
    return
  }
  var phone = scene.getObjectByName("电话").text
  if(phone.length < 3 ){
    // Mugine.Utils.Toast.info("未输入电话", {type:'info'});
    Swal.fire({
      icon: 'error',
      text: "未输入电话",
      confirmButtonText: "好的"
    })
    return
  }
  var address = scene.getObjectByName("地址").text
  if(address.length < 1 ){
    Swal.fire({
      icon: 'error',
      text: "未输入地址",
      confirmButtonText: "好的"
    })
    return
  }
  // 提交个人信息 
  fetch(domain+"/v3/mugeda/form/v3/userinfo?name="+name+"&phone="+phone+"&address="+address,{
    headers: {
      'wechat_jwt_token': wechat_jwt_token,
    },
    method: 'PUT', // *GET, POST, PUT, DELETE, etc.
    mode: 'cors', // no-cors, cors, *same-origin
    redirect: 'follow', // manual, *follow, error
    referrer: 'no-referrer', // *client, no-referrer
  })
  .then((res)=>res.json())
  .then((res)=>{
    if(res.code == 1){
      // scene.gotoAndPause(1, 5);
      userinfo = res.data

      Swal.fire({
        icon: 'success',
        title: res.msg,
        confirmButtonText: "好的"
      })
      console.log(res.data);
    }else{
      Swal.fire({
        icon: 'error',
        title: res.msg,
        text: res.err,
        confirmButtonText: "好的"
      })
      Mugine.Utils.Toast.info( res.msg+"err_msg:"+res.err, {type:'info'});
    }
  })
  .catch((err)=>{
    Swal.fire({
      icon: 'error',
      text: err,
      confirmButtonText: "好的"
    })
  });
}




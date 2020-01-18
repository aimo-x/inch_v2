// 以来 导入主题 https://cdn.jsdelivr.net/npm/@sweetalert2/theme-borderless/borderless.css
function dynamicLoadCss(url) {
  var head = document.getElementsByTagName('head')[0];
  var link = document.createElement('link');
  link.type = 'text/css';
  link.rel = 'stylesheet';
  link.href = url;
  head.appendChild(link);
}
const h5_link = "https://6.u.mgd5.com/c/5z0l/bnvn/index.html"
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
        if(decodeURIComponent(temp[1]) == ""){
          return false;
        }else{
          return decodeURIComponent(temp[1]);
        }
        
    }    
  }
  return false;
};
wechat_jwt_token = getCookie("wechat_jwt_token_fv3");

wechat_oauth = () => {
    // var reuri = window.location.origin+window.location.pathname;
    var reuri = window.location.href;
    fetch(domain+"/v3/mugeda/form/v3/oauth/wechat?state="+encodeURIComponent(reuri))
    .then((res)=>res.json())
    .then((res)=>{
        if(res.code == 1){
            window.location.href = res.data;
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
userinfo="";
user_openid = ""
// 获取用户信息
GetUserInfo = () => {
  var scene = mugeda.scene;
  window.s = scene
  /*
  var head_img = scene.getObjectByName("头像")
  var nick_name = scene.getObjectByName("微信昵称")
  */
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
        user_openid = res.data.open_id
        /*
        head_img.src = "https" + res.data.head_img.substring(4)
        nick_name.text = Base64.decode(res.data.nick_name)
        */
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
            setCookie("wechat_jwt_token_fv3",token,1/48);
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



var bless_id = search('bless_id')
var bless_receive_id = search('bless_receive_id')
var mugeda_form_v3_bless = ""

// 进行助力
post_bless_receive_invite = () => {
  // bless/receive/invite
  fetch(domain+"/v3/mugeda/form/v3/bless/receive/invite?bless_receive_id="+bless_receive_id,{
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
      // 查询头像进行
      get_bless_receive_headimg(res.data.invite)
      Swal.fire({
        icon: 'success',
        title: "助力成功",
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
      //Mugine.Utils.Toast.info( res.msg+"err_msg:"+res.err, {type:'info'});
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
// 木疙瘩 api
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
  // 检查是否登陆，是否需要助力
  if(bless_receive_id && wechat_jwt_token){
    // 执行助力接口 
    post_bless_receive_invite()
    scene.gotoAndPause(0, 1);
  }
  // 检查是否存在此ID 存在即跳转页面
  if(bless_id && wechat_jwt_token){
    // 跳转到邀请页面
    scene.gotoAndPause(0, 1);
  }

});


// res.data.mugeda_form_v3_bless

// 拉取已经助力的头像
get_bless_receive_headimg = (invite) => {
  fetch(domain+"/v3/mugeda/form/v3/userinfo/arr?open_id_arr="+invite,{
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
      var head_img = []
      for(var i = 0; i<res.data.length;i++){
        var n = i + 1
        head_img[i] = scene.getObjectByName("好友"+n)
        head_img[i].src = "https" + res.data.head_img.substring(4)
      }
      if(res.data.length > 3){
        // 集齐4个好友
        scene.gotoAndPause(1, 1);
      }
      console.log(res.data);
    }else{
      Swal.fire({
        icon: 'error',
        title: res.msg,
        text: res.err,
        confirmButtonText: "好的"
      })
      //Mugine.Utils.Toast.info( res.msg+"err_msg:"+res.err, {type:'info'});
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
// 查看祝福 执行此函数 
view_mugeda_form_v3_bless = () => {
  if(mugeda_form_v3_bless == ""){
    Swal.fire({
      // icon: 'error',
      text: "暂时无法查看祝福",
      confirmButtonText: "好的"
    })
  }else{
    Swal.fire({
      // icon: 'success',
      text: mugeda_form_v3_bless.content,
      confirmButtonText: "好的"
    })
  }
}
// fq_invite 发起邀请
fq_invite = () => {
  // 设置分享
  defineWechatParameters({
    url_callback: function(){
        return h5_link+"?bless_id="+bless_id
    }
  });
  Swal.fire({
    icon: 'success',
    text: "点击右上角分享给好友助力",
    confirmButtonText: "好的"
  })
}

// 接收/查询祝福语 不存在则自动创建/满足条件即可查看/祝福语
get_bless_receive = () => {

  if(!bless_id){
    
    Swal.fire({
      icon: 'error',
      text: "您还没有收到祝福",
      confirmButtonText: "好的"
    })
    return
  }
  fetch(domain+"/v3/mugeda/form/v3/bless/receive?bless_id="+bless_id,{
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
      // scene.gotoAndPause(1, 5);
      var invite = res.data.mugeda_form_v3_bless_receive.invite
      mugeda_form_v3_bless = res.data.mugeda_form_v3_bless
      // var open_id_arr = invite.split(",")
      if(invite == ""){
        console.log("啥也不用干")
        return
      }
      // 继续拉取已经助力的头像
      get_bless_receive_headimg(invite)
      console.log(res.data);
    }else{
      Swal.fire({
        icon: 'error',
        title: res.msg,
        text: res.err,
        confirmButtonText: "好的"
      })
      //Mugine.Utils.Toast.info( res.msg+"err_msg:"+res.err, {type:'info'});
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

// 生成带参数海报二维码
get_qrcode = (content) => {
  fetch("https://iuu.pub/api/qr?url="+content)
    .then((res)=>res.json())
    .then((res)=>{
        if(res.code == 1){
          var qr = scene.getObjectByName("二维码");
          qr.src = "data:image/png;base64,"+res.data
          // scene.gotoPage(5, options);
        }else{
          // Mugine.Utils.Toast.info(res.msg+",error:"+res.err, {type:'info'});
          Swal.fire({
            icon: 'error',
            title: res.msg,
            text: res.err,
            confirmButtonText: "好的"
          })
        }
    })
    .catch((err)=>{
        // Mugine.Utils.Toast.info(err, {type:'info'});
        Swal.fire({
          icon: 'error',
          text: err,
          confirmButtonText: "好的"
        })
    })
}
// 创建祝福语 并加入阵营
post_bless_content = () => {
  var camp_id = scene.getObjectByName("阵营ID").text
  var content = scene.getObjectByName("祝福").text
  fetch(domain+"/v3/mugeda/form/v3/bless?content="+content+"&camp_id="+camp_id,{
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
      // scene.gotoAndPause(1, 5);
      // 生成带参数海报二维码
      get_qrcode(h5_link+"?"+res.data.ID)
      scene.gotoAndPause(0, 4);
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
      //Mugine.Utils.Toast.info( res.msg+"err_msg:"+res.err, {type:'info'});
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

      Swal.fire({
        icon: 'success',
        title: "提交成功",
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
      //Mugine.Utils.Toast.info( res.msg+"err_msg:"+res.err, {type:'info'});
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




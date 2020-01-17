const domain = "face.inch.online";
// options的取值控制翻页效果，形式如下：
const options = {
    transition: "", // 过渡效果
    exit: "", // 退出效果
    duration: 1, // 过渡时间
};
// 图层顺序调整

// 开始融合
mugeda_face_merge = () => {
    window.scene = mugeda.scene;
    scene.gotoPage(6, options); // 调整到加载页
    var sex = scene.getObjectByName("性别").text
    var sym = scene.getObjectByName(sex+"照片集").scene;
    var template_id = Math.round((sym.currentId+1)/5) + 1
    console.log("f",template_id,"sex",sex,"sym.currentId",sym.currentId)
    var image_target  = scene.getObjectByName("上传照片").src;
    var image_template = "";
    if(sex == "女"){
        // image_template = "https://www.inch.online/v2/usr/face/20200110/NV"+template_id+".jpg";
        image_template = scene.getObjectByName("NV"+template_id).src;

    }else{
        // image_template = "https://www.inch.online/v2/usr/face/20200110/NA"+template_id+".jpg";
        image_template = scene.getObjectByName("NA"+template_id).src;
    }
    window.merged_show = scene.getObjectByName("容器");
    var JSONdata = {
        image_target: "https:"+image_target,
        image_template: image_template,
    };
    fetch('https://'+domain+'/api/mugeda/face/merge',{
        headers: { 
            "Content-Type": "application/json"
        },
        method: 'POST',
        body: JSON.stringify(JSONdata)
    })
    .then(function(response) {
        return response.json();
    })
    .then(function(res) {
        if(res.code == 1 ){
           // 提交成功
           merged_show.src = "data:image/jpg;base64,"+res.data;
           // scene.gotoAndPause(2, 5);
           scene.gotoPage(3, options);
        }else{
            swal("发生错误", res.msg, "error");
            scene.gotoPage(1, options);
        }
    })
    .catch((err)=>{
        swal("发生错误", err, "error");
        scene.gotoPage(1, options);
    });
};
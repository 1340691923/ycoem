<template>
    <div id="upload">
      <!--elementui的上传文件的upload组件-->
      <el-upload class="upload-demo"
                 ref="upload"
                 :headers="getHeader"
                 :file-list="fileList"
                 :action="uploadUrl"
                 :limit="limit"
                 :on-success="UploadSuccess"
                 :on-preview="handlePictureCardPreview"
                 :on-exceed="exceedHandle"
                 :auto-upload="true"
                 :on-remove="removeHandle"
                 :multiple='true'>
        <el-button size="small" type="primary" icon="el-icon-upload">点击上传</el-button>
      </el-upload>
      <!--展示选中文件的区域-->
      <el-dialog :visible.sync="dialogImgVisible" style="position: absolute;top: 0px;z-index: 25555">
        <img width="100%"
             :src="dialogImageUrl"
             alt="">
      </el-dialog>
    </div>

</template>

<script>
  import { getToken } from '@/utils/auth'
  export default {
    name: 'index',
    props:[
      'limit',
      'data',
      "index"
    ],
    data(){
      return {
        dialogImageUrl: '',
        dialogImgVisible: false,
        fileList: this.data,
        param: '', // 表单要提交的参数
        src: '' // 展示文件的地址
      }
    },
    computed:{
      getHeader(){
        let header = {}
        header["x-token"] = getToken()
        return header
      },
      uploadUrl(){
        return process.env.VUE_APP_BASE_API+"/api/upload/zip"
      }
    },
    methods: {
      handleRemoveImage(file) {
        //缩略图列表跟上传文件列表并不一致
        this.$refs.upload.handleRemove(file); //删除缩略图
      },
      // 3，点击文件列表中已上传的文件时的钩子
      handlePictureCardPreview (file) {
        this.$emit('handlePictureCardPreview',file)
      },
      UploadSuccess(res, file, fileList){
        if(res.code !=0){
          this.$message({
            type: 'error',
            message: res.msg,
            duration:8000,
            showClose:true
          })
          this.handleRemoveImage(file)
          return
        }
        this.$message({
          type: 'success',
          message: res.msg
        })

        for(let index in fileList){
          if(fileList[index].uid){
            if(fileList[index].uid == file.uid){
              fileList[index].url = res.data.url
            }
          }
        }
        console.log("fileList",fileList)
        this.data = fileList
        this.$emit('changeFileList',this.data)
      },
      // 5设置超过9张图给与提示
      exceedHandle () {
        this.$message({
          type: 'error',
          message: "您现在选择已超过"+this.limit+"个文件，请重新选择"
        })
      },
      removeHandle(file, fileList){
        this.fileList = fileList
        this.$emit('changeFileList',this.fileList,this.index)
      }
    }
  }
</script>

<style scoped>

</style>

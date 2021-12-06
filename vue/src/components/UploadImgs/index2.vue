<template>
  <div class="app-container">
<el-card class="box-card">
    <div id="upload">
      <!--elementui的上传图片的upload组件-->
      <el-upload class="upload-demo"
                 ref="upload"
                 :file-list="data"
                 list-type="picture-card"
                 :action="uploadUrl"
                 :limit="limit"
                 :on-success="UploadSuccess"

                 :on-preview="handlePictureCardPreview"
                 :on-exceed="exceedHandle"
                 :auto-upload="true"
                 :multiple='true'>
        <i class="el-icon-plus"></i>
      </el-upload>
      <!--展示选中图片的区域-->
      <el-dialog :visible.sync="dialogVisible">
        <img width="100%"
             :src="dialogImageUrl"
             alt="">
      </el-dialog>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'index',
    props:[
      'limit',
      'data'
    ],
    data(){
      return {
        dialogImageUrl: '',
        dialogVisible: false,
        fileList: [],
        param: '', // 表单要提交的参数
        src: '' // 展示图片的地址
      }
    },
    computed:{
      uploadUrl(){
        return process.env.VUE_APP_BASE_API+"/api/upload/img"
      }
    },
    methods: {
      // 3，点击文件列表中已上传的文件时的钩子
      handlePictureCardPreview (file) {
        if (file.url){
          this.dialogImageUrl = file.url
        }else{
          this.dialogImageUrl = file
        }
        this.dialogImgVisible = true
      },
      UploadSuccess(response, file, fileList){
        this.data = fileList
        this.$emit('changeFileList',this.data)
      },
      // 5设置超过9张图给与提示
      exceedHandle () {
        this.$message({
          type: 'error',
          message: "您现在选择已超过"+this.limit+"张图，请重新选择"
        })
      }
    }
  }
</script>

<style scoped>

</style>

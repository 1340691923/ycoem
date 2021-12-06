<template>
  <div class="app-container">
<el-card class="box-card">
    <el-form :model="form" ref="form" label-width="100px" class="demo-dynamic">
      <el-form-item
        prop="copyright"
        label="版权所属"
      >
        <el-input type="textarea"
                  autosize v-model="form.copyright" style="width:300px"></el-input>
      </el-form-item>
      <el-form-item
        prop="icp"
        label="icp"
      >
        <el-input style="width:300px" type="textarea"
                  autosize v-model="form.icp"></el-input>
      </el-form-item>
      <el-form-item
        prop="record_no"
        label="备案号"
      >
        <el-input style="width:300px" type="textarea"
                  autosize v-model="form.record_no"></el-input>
      </el-form-item>
      <el-form-item
        prop="record_info"
        label="备案信息"
      >
        <el-input style="width:300px" type="textarea"
                  autosize v-model="form.record_info"></el-input>
      </el-form-item>
      <el-form-item
        prop="record_info"
        label="公众号二维码"
      >
        <upload-imgs   @changeFileList="changeWXList" :data="wxList" :limit="limit"></upload-imgs>
      </el-form-item>
        <el-form-item
          prop="record_info"
          label="APP二维码"
        >
          <upload-imgs   @changeFileList="changeAPPList" :data="appList" :limit="limit"></upload-imgs>
        </el-form-item>
            <el-form-item>
        <el-button type="primary" @click="submitForm('form')" icon="el-icon-check">提交</el-button>
        <el-button @click="resetForm()" icon="el-icon-delete">重置</el-button>
      </el-form-item>
    </el-form>

</el-card></div>
</template>

<script>
  import {SetWebinfoAction,GetWebinfoAction} from "../../api/seo"
  import {getImageName} from "../../utils/upload"
  import {deepClone} from "../../utils"
  export default {
    name: 'index',
    components:{
      UploadImgs :()=>import("@/components/UploadImgs"),
    },
    data(){
      return {
        limit:1,
        form:{
          copyright:"",
          icp:"",
          record_no:"",
          record_info:"",
          wx_img:"",
          app_img:""
        },
        wxList:[],
        appList:[]
      }
    },
    mounted(){
      this.getSeo()
    },
    methods:{
      changeWXList(v){
        console.log("v",v)
        this.wxList = v
      },
      changeAPPList(v){
        this.appList = v
      },
      getSeo(){
        GetWebinfoAction().then(res=>{
          if (res.code!=0){
            this.$message({
              message: res.msg,
              type: 'error'
            })
            return
          }
          this.$message({
            message: res.msg,
            type: 'success'
          })
          this.form = res.data
          if(this.form.wx_img!=""){
            this.wxList.push({"url":this.form.wx_img})
          }
          if(this.form.app_img!=""){
            this.appList.push({"url":this.form.app_img})
          }
        })
          .catch(err=>{
            console.log(err)
            this.$message({
              message: "网络异常",
              type: 'error'
            })
          })
      },
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.form.app_img = getImageName(this.appList,0)
            this.form.wx_img = getImageName(this.wxList,0)
            SetWebinfoAction(this.form).then(res=>{
              if (res.code != 0){
                this.$message({
                  message: res.msg,
                  type: 'error'
                })
                return
              }
              this.resetForm()
              this.$message({
                message: res.msg,
                type: 'success'
              })
              this.getSeo()
            })
              .catch(err=>{
                console.log(err)
                this.$message({
                  message: "网络异常",
                  type: 'error'
                })
              })
          } else {
            this.$message({
              message: '请检查表单',
              type: 'error'
            })
            return false;
          }
        });
      },
      resetForm(formName) {
        this.form = deepClone({})
        this.appList = []
        this.wxList = []
      },
    }
  }
</script>

<style scoped>

</style>

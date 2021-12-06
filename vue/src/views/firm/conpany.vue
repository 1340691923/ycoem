<template>
  <div>
    <div class="app-container">
<el-card class="box-card">
      <el-form :model="dynamicValidateForm" ref="dynamicValidateForm" label-width="100px" class="demo-dynamic">

        <el-form-item
          prop="content"
          label="内容"
        >
          <el-input style="width:300px" type="textarea"
                    autosize v-model="dynamicValidateForm.content"></el-input>
        </el-form-item>

        <el-form-item
          v-for="(keywords, index) in dynamicValidateForm.list"
          :label="'企业文化' + Number(index+1)"
          :key="keywords.key"
        >
          <el-input v-model="keywords.value"  style="width:300px"></el-input><el-button @click.prevent="removeKeyword(keywords)" icon="el-icon-delete">删除</el-button>
        </el-form-item>

        <el-form-item
          prop="img"
          label="图片"
        >
          <upload-imgs   @changeFileList="changeWXList" :data="wxList" :limit="limit"></upload-imgs>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="submitForm()" icon="el-icon-check">提交</el-button>
          <el-button @click="addKeyword" icon="el-icon-plus">新增企业文化</el-button>
          <el-button @click="resetForm('dynamicValidateForm')" icon="el-icon-delete">重置</el-button>
        </el-form-item>
      </el-form>
</el-card></div>
  </div>
</template>

<script>
  import {getImageName} from "../../utils/upload"
  import {SetIntroduceAction,GetIntroduceAction} from "../../api/firm"
  export default {
    name: 'index',
    data(){
      return {
        form:{
          img:"",
          content:"",
          list:[]
        },
        dynamicValidateForm: {
          list: [{
            value: ''
          }],
          img: '',
          content:"",
        },
        wxList:[],
        limit:1
      }
    },
    components:{
      UploadImgs :()=>import("@/components/UploadImgs"),
    },
    mounted(){
      this.getSeo()
    },
    methods:{
      changeWXList(v){
        this.wxList = v
      },
      getSeo(){
        GetIntroduceAction().then(res=>{
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
          this.dynamicValidateForm.content = res.data.content
          this.dynamicValidateForm.list = []
          let arr = res.data.list
          for (let i in arr){
            this.dynamicValidateForm.list.push({
              value: arr[i],
              key:i
            });
          }
          this.dynamicValidateForm.img = res.data.img

          if(this.dynamicValidateForm.img != ""){
            this.wxList.push({"url":this.dynamicValidateForm.img})
          }
          console.log(res,"res")
        })
          .catch(err=>{
            console.log(err)
            this.$message({
              message: "网络异常",
              type: 'error'
            })
          })
      },
      submitForm() {

        this.form.img = getImageName(this.wxList,0)
        this.form.content = this.dynamicValidateForm.content

            let keywordsArr = []
            for (let v of this.dynamicValidateForm.list) {
              keywordsArr.push(v["value"])
            }
            this.form.list = keywordsArr
            SetIntroduceAction(this.form).then(res=>{
              if (res.code != 0){
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
              this.wxList = []
              this.getSeo()
            })
              .catch(err=>{
                console.log(err)
                this.$message({
                  message: "网络异常",
                  type: 'error'
                })
              })

      },
      resetForm(formName) {
        this.wxList = []
        this.dynamicValidateForm.img = ""
        this.dynamicValidateForm.content = ""
        this.dynamicValidateForm.list = []
        this.dynamicValidateForm.list.push({
          value: '',
          key: Date.now()
        });
      },
      removeKeyword(item) {
        var index = this.dynamicValidateForm.list.indexOf(item)
        if (index !== -1) {
          this.dynamicValidateForm.list.splice(index, 1)
        }
      },
      addKeyword() {
        console.log(this.dynamicValidateForm)
        this.dynamicValidateForm.list.push({
          value: '',
          key: Date.now()
        });
      }
    }
  }
</script>

<style scoped>

</style>

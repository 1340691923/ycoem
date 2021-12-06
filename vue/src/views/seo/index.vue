<template>
    <div>
      <div class="app-container">
<el-card class="box-card">
        <el-form :model="dynamicValidateForm" ref="dynamicValidateForm" label-width="100px" class="demo-dynamic">
          <el-form-item
            prop="email"
            label="title"
          >
          <el-input type="textarea"
                    autosize v-model="dynamicValidateForm.title" style="width:300px"></el-input>
          </el-form-item>
          <el-form-item
            prop="email"
            label="描述"
          >
            <el-input style="width:300px" type="textarea"
                      autosize v-model="dynamicValidateForm.description"></el-input>
          </el-form-item>
          <el-form-item
            v-for="(keywords, index) in dynamicValidateForm.keywords"
            :label="'关键词' + Number(index+1)"
            :key="keywords.key"
            :prop="'keywords.' + index + '.value'"
            :rules="{
      required: true, message: '关键词不能为空', trigger: 'blur'
    }"
          >
            <el-input v-model="keywords.value"  style="width:300px"></el-input><el-button @click.prevent="removeKeyword(keywords)" icon="el-icon-delete" >删除</el-button>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm('dynamicValidateForm')" icon="el-icon-check">提交</el-button>
            <el-button @click="addKeyword" icon="el-icon-plus">新增关键词</el-button>
            <el-button @click="resetForm('dynamicValidateForm')" icon="el-icon-delete">重置</el-button>
          </el-form-item>
        </el-form>
</el-card></div>
    </div>
</template>

<script>
  import {create,get} from "../../api/seo"
  export default {
    name: 'index',
    data(){
      return {
        form:{
          title:"",
          description:"",
          keywords:""
        },
        dynamicValidateForm: {
          keywords: [{
            value: ''
          }],
          title: '',
          description:"",
        }
      }
    },
    mounted(){
      this.getSeo()
    },
    methods:{
      getSeo(){
        get().then(res=>{
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
          this.dynamicValidateForm.title = res.data.title
          this.dynamicValidateForm.description = res.data.description
          this.dynamicValidateForm.keywords = []
          let arr = res.data.keywords.split(",")
          for (let i in arr){
            this.dynamicValidateForm.keywords.push({
              value: arr[i],
              key:i
          });
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
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.form.title = this.dynamicValidateForm.title
            this.form.description = this.dynamicValidateForm.description
            let keywordsArr = []
            for (let v of this.dynamicValidateForm.keywords) {
              keywordsArr.push(v["value"])
            }
            this.form.keywords = keywordsArr.join(",")
            create(this.form).then(res=>{
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
        this.dynamicValidateForm.title = ""
        this.dynamicValidateForm.description = ""
        this.dynamicValidateForm.keywords = []
        this.dynamicValidateForm.keywords.push({
          value: '',
          key: Date.now()
        });
      },
      removeKeyword(item) {
        var index = this.dynamicValidateForm.keywords.indexOf(item)
        if (index !== -1) {
          this.dynamicValidateForm.keywords.splice(index, 1)
        }
      },
      addKeyword() {
        this.dynamicValidateForm.keywords.push({
          value: '',
          key: Date.now()
        });
      }
    }
  }
</script>

<style scoped>

</style>

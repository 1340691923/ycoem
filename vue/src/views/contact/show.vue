<template>
  <div>
    <div class="app-container">
<el-card class="box-card">
      <el-form :model="dynamicValidateForm" ref="dynamicValidateForm" label-width="100px" class="demo-dynamic">
        <el-form-item
          prop="email"
          label="总部地址"
        >
          <el-input type="textarea"
                    autosize v-model="dynamicValidateForm.address" style="width:300px"></el-input>
        </el-form-item>
        <el-form-item
          v-for="(dept_infos, index) in dynamicValidateForm.dept_infos"
          :label="'展示公司' + Number(index+1)"
          :key="dept_infos.key"
          :prop="'dept_infos.' + index + '.value'"
          :rules="{
      required: true, message: '您所展示的公司不能为空', trigger: 'blur'
    }"
        >
          <el-input v-model="dept_infos.value"  style="width:300px"></el-input><el-button icon="el-icon-close" @click.prevent="removeKeyword(dept_infos)" >删除</el-button>
        </el-form-item>
        <el-form-item>
          <el-button icon="el-icon-check" type="primary" @click="submitForm('dynamicValidateForm')">提交</el-button>
          <el-button icon="el-icon-plus" @click="addKeyword">新增展示公司</el-button>
          <el-button icon="el-icon-delete" @click="resetForm('dynamicValidateForm')">重置</el-button>
        </el-form-item>
      </el-form>
</el-card></div>
  </div>
</template>

<script>
  import {GetContactAction,SetContactAction} from "../../api/concat"
  export default {
    name: 'index',
    data(){
      return {
        form:{
          address:"",
          dept_infos:""
        },
        dynamicValidateForm: {
          dept_infos: [{
            value: ''
          }],
          address: '',
        }
      }
    },
    mounted(){
      this.getSeo()
    },
    methods:{
      getSeo(){
        GetContactAction().then(res=>{
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
          this.dynamicValidateForm.address = res.data.address
          this.dynamicValidateForm.dept_infos = []
          let arr = res.data.dept_infos.split(",")
          for (let i in arr){
            this.dynamicValidateForm.dept_infos.push({
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
            this.form.address = this.dynamicValidateForm.address
            let dept_infosArr = []
            for (let v of this.dynamicValidateForm.dept_infos) {
              dept_infosArr.push(v["value"])
            }
            this.form.dept_infos = dept_infosArr.join(",")
            SetContactAction(this.form).then(res=>{
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
        this.dynamicValidateForm.address = ""
        this.dynamicValidateForm.dept_infos = []
        this.dynamicValidateForm.dept_infos.push({
          value: '',
          key: Date.now()
        });
      },
      removeKeyword(item) {
        var index = this.dynamicValidateForm.dept_infos.indexOf(item)
        if (index !== -1) {
          this.dynamicValidateForm.dept_infos.splice(index, 1)
        }
      },
      addKeyword() {
        this.dynamicValidateForm.dept_infos.push({
          value: '',
          key: Date.now()
        });
      }
    }
  }
</script>

<style scoped>

</style>

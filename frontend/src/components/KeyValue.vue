<template>
  <main>
    <el-form :model="form">
      <el-form-item label="键">
        <el-input v-model="form.key" disabled placeholder="" />
      </el-form-item>
      <el-row>
        <el-col :span="11">
          <el-form-item label="过期时间">
            <el-input v-model="form.ttl" placeholder="" />
          </el-form-item>
        </el-col>
        <el-col :span="2"></el-col>
        <el-col :span="11">
          <el-form-item label="数据类型">
            <el-input v-model="form.type" disabled placeholder="" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item label="值">
        <el-input type="textarea" v-model="form.value" placeholder="" />
      </el-form-item>
    </el-form>
  </main>
</template>

<script setup>
import {ref, watch} from 'vue'
import {GetKeyValue} from "../../wailsjs/go/main/App.js";
import {ElNotification} from "element-plus";
let props = defineProps(['keyDB', 'keyConnIdentity', 'keyKey'])
let form = ref({})

watch(()=>props.keyKey, () => {
  getTheValue()
})

function getTheValue() {
  GetKeyValue({conn_identity:props.keyConnIdentity, db: props.keyDB, key: props.keyKey}).then(res => {
    if (res.code !== 200) {
      ElNotification({
        title:res.msg,
        type: "error",
      })
      return
    }
    form.value = res.data
    form.value.key = props.keyKey
  })
}
</script>

<style scoped>

</style>
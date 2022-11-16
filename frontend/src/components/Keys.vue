<template>
  <main>
    <el-form :inline="true" :model="form" class="demo-form-inline">
      <el-form-item label="">
        <el-input v-model="form.keyword" placeholder="请输入键的信息" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="getKeyList">查询</el-button>
      </el-form-item>
    </el-form>
    <div v-for="item in keys" class="my-item">{{item}}</div>
  </main>
</template>

<script setup>
import {KeyList} from "../../wailsjs/go/main/App.js";
import {reactive, ref, watch} from 'vue'
import {ElNotification} from "element-plus";

const form = reactive({
  keyword: '',
})

let props = defineProps(['keyDB', 'keyConnIdentity'])
let keys = ref()

watch(props, ()=> {
  getKeyList()
})

function getKeyList() {
  let data = {
    conn_identity: props.keyConnIdentity,
    db: props.keyDB,
    keyword: form.keyword
  }
  KeyList(data).then(res => {
    if (res.code !== 200) {
      ElNotification({
        title:res.msg,
        type: "error",
      })
      return
    }
    keys.value = res.data
  })
}
</script>

<style scoped>

</style>
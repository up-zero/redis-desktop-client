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
    <div v-for="item in keys" @click="selectKeyKey(item)">
      <div v-if="item === selectKey" class="my-select-item">{{item}}</div>
      <div v-else class="my-item">{{item}}</div>
    </div>
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
let emits = defineEmits(['emit-select-key'])
let keys = ref()
let selectKey = ref()

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

function selectKeyKey(key) {
  selectKey.value = key
  emits('emit-select-key', key)
}
</script>

<style scoped>

</style>
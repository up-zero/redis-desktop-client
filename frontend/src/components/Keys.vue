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
      <div v-if="item === selectKey" class="item key-select-item">
        <div style="padding: 5px 12px">{{item}}</div>
        <el-popconfirm title="确认删除?" @confirm="deleteKey(item)">
          <template #reference>
            <el-button text type="danger" @click.stop>删除</el-button>
          </template>
        </el-popconfirm>
      </div>
      <div v-else class="item key-item">
        <div style="padding: 5px 12px">{{item}}</div>
        <el-popconfirm title="确认删除?" @confirm="deleteKey(item)">
          <template #reference>
            <el-button text type="danger" @click.stop>删除</el-button>
          </template>
        </el-popconfirm>
      </div>
    </div>
  </main>
</template>

<script setup>
import {ConnectionDelete, DeleteKeyValue, KeyList} from "../../wailsjs/go/main/App.js";
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

function deleteKey(key) {
  DeleteKeyValue({conn_identity: props.keyConnIdentity, db: props.keyDB, key: key}).then(res => {
    if (res.code !== 200) {
      ElNotification({
        title:res.msg,
        type: "error",
      })
      return
    }
    ElNotification({
      title:res.msg,
      type: "success",
    })
    getKeyList()
  })
}
</script>

<style scoped>
.key-item {
  color: #409eff;
  background-color: #ecf5ff;
  margin-bottom: 5px;
}
.key-select-item {
  color: #67c23a;
  background-color: #f0f9eb;
  margin-bottom: 5px;
}
</style>
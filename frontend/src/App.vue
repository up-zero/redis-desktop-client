<script setup>
import ConnectionList from "./components/ConnectionList.vue";
import ConnectionManage from "./components/ConnectionManage.vue";
import {ref} from "vue";
import {DbList, GetKeyValue, KeyList} from "../wailsjs/go/main/App.js";
import Keys from "./components/Keys.vue";
import KeyValue from "./components/KeyValue.vue";

let flushFlag = ref(true)
let keyDB = ref()
let keyConnIdentity = ref()
let keyKey = ref()

function flushConnectionList() {
  flushFlag.value = !flushFlag.value
}

// 选中数据库
function selectDB (db, connIdentity) {
  keyDB.value = db
  keyConnIdentity.value = connIdentity
}

// 选中键
function selectKey(key) {
  keyKey.value = key
}
</script>

<template>
  <el-row>
    <el-col :span="5" style="height: 100vh; padding: 12px">
      <div style="margin-bottom: 12px">
        <ConnectionManage title="新建连接" btn-type="primary" @emit-connection-list="flushConnectionList"/>
      </div>
      <ConnectionList @emit-select-db="selectDB" :flush="flushFlag"/>
    </el-col>
    <el-col :span="7" style="padding: 12px">
      <Keys :keyDB="keyDB" :keyConnIdentity="keyConnIdentity" @emit-select-key="selectKey"/>
    </el-col>
    <el-col :span="12" style="padding: 12px">
      <KeyValue :keyDB="keyDB" :keyConnIdentity="keyConnIdentity" :keyKey="keyKey" />
    </el-col>
  </el-row>
</template>

<style>
</style>

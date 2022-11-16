<script setup>
import ConnectionList from "./components/ConnectionList.vue";
import ConnectionManage from "./components/ConnectionManage.vue";
import {ref} from "vue";
import {DbList, KeyList} from "../wailsjs/go/main/App.js";
import Keys from "./components/Keys.vue";

let flushFlag = ref(true)
let keyDB = ref()
let keyConnIdentity = ref()
let keyList = ref()

function flushConnectionList() {
  flushFlag.value = !flushFlag.value
}

// 选中数据库
function selectDB (db, connIdentity) {
  keyDB.value = db
  keyConnIdentity.value = connIdentity
}

KeyList({conn_identity:"cc69d2e0-80e4-40ed-96a3-8706403b4c7c", db:0, keyword: "name"}).then(res => {
  keyList.value = res
})
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
      <Keys :keyDB="keyDB" :keyConnIdentity="keyConnIdentity"/>
    </el-col>
    <el-col :span="12">
      dbList ==> {{keyList}}
    </el-col>
  </el-row>
<!--  <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>-->
<!--  <HelloWorld/>-->
</template>

<style>
</style>

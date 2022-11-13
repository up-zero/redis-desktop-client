<script setup>
import HelloWorld from './components/HelloWorld.vue'
import ConnectionList from "./components/ConnectionList.vue";
import ConnectionManage from "./components/ConnectionManage.vue";
import {ref} from "vue";
import {DbList} from "../wailsjs/go/main/App.js";

let flushFlag = ref(true)

let dbList = ref()

function flushConnectionList() {
  flushFlag.value = !flushFlag.value
}

DbList("cc69d2e0-80e4-40ed-96a3-8706403b4c7c").then(res => {
  dbList.value = res
})
</script>

<template>
  <el-row>
    <el-col :span="6" style="height: 100vh; padding: 12px">
      <div style="margin-bottom: 12px">
        <ConnectionManage title="新建连接" btn-type="primary" @emit-connection-list="flushConnectionList"/>
      </div>
      <ConnectionList :flush="flushFlag"/>
    </el-col>
    <el-col :span="18">
      dbList ==> {{dbList}}
    </el-col>
  </el-row>
<!--  <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>-->
<!--  <HelloWorld/>-->
</template>

<style>
</style>

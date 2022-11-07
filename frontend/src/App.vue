<script setup>
import HelloWorld from './components/HelloWorld.vue'
import ConnectionList from "./components/ConnectionList.vue";
import ConnectionManage from "./components/ConnectionManage.vue";
import {ref} from "vue";

let flushFlag = ref(true)

function flushConnectionList() {
  flushFlag.value = !flushFlag.value
}
</script>

<template>
  <el-row>
    <el-col :span="6" style="height: 100vh; margin: 12px">
      <div style="margin-bottom: 12px">
        <ConnectionManage title="新建连接" @emit-connection-list="flushConnectionList"/>
      </div>
      <ConnectionList :flush="flushFlag"/>
    </el-col>
  </el-row>
<!--  <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>-->
<!--  <HelloWorld/>-->
</template>

<style>
</style>

<script setup>
import HelloWorld from './components/HelloWorld.vue'
import ConnectionList from "./components/ConnectionList.vue";</script>

<template>
  <el-row>
    <el-col :span="6" style="height: 100vh">
      <ConnectionList />
    </el-col>
  </el-row>
<!--  <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>-->
<!--  <HelloWorld/>-->
</template>

<style>
</style>

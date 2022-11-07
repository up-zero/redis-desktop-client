<template>
  <main>
    <el-button type="primary" @click="dialogVisible = true">{{title}}</el-button>
    <el-dialog
        v-model="dialogVisible"
        :title="title"
        width="60%"
    >
      <el-form :model="form" label-width="120px">
        <el-form-item label="连接地址">
          <el-input placeholder="请输入连接地址" v-model="form.addr" />
        </el-form-item>
        <el-form-item label="连接名称">
          <el-input placeholder="请输入连接名称" v-model="form.name" />
        </el-form-item>
        <el-form-item label="端口号">
          <el-input placeholder="请输入端口号" v-model="form.port" />
        </el-form-item>
        <el-form-item label="用户名">
          <el-input placeholder="请输入用户名" v-model="form.username" />
        </el-form-item>
        <el-form-item label="密码">
          <el-input placeholder="请输入密码" type="password" v-model="form.password" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="createConnection">创建</el-button>
          <el-button @click="dialogVisible = false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

  </main>
</template>

<script setup>
import {ref, reactive} from "vue";
import {ConnectionCreate} from "../../wailsjs/go/main/App.js";
import {ElNotification} from "element-plus";
defineProps(['title'])
const dialogVisible = ref(false)
const emits = defineEmits(['emit-connection-list'])
const form = reactive({})

function createConnection() {
  ConnectionCreate(form).then(res => {
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
    // 获取新的连接地址
    emits('emit-connection-list')
  })
}
</script>

<style scoped>

</style>
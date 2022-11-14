<template>
  <main>
    <div class="demo-collapse">
      <el-collapse accordion>
        <el-collapse-item v-for="item in list" :name="item.identity" @click="getInfo(item.identity)">
          <template #title>
            <div class="item">
              <div>
                {{ item.name }}
              </div>
              <div style="display: flex">
                <ConnectionManage @click.stop title="编辑" btn-type="text" :data="item" @emit-connection-list="connectionList"/>
                <el-popconfirm title="确认删除?" @confirm="connectionDelete(item.identity)">
                  <template #reference>
                    <el-button link type="danger" @click.stop>删除</el-button>
                  </template>
                </el-popconfirm>
              </div>
            </div>
          </template>
          <div v-for="db in infoDbList" class="db-item">
            {{db.key}} ( {{db.number}} )
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
  </main>
</template>

<script setup>
import {ref, watch} from "vue";
import {ConnectionDelete, ConnectionList, DbList} from "../../wailsjs/go/main/App.js";
import {ElNotification} from "element-plus"
import ConnectionManage from "./ConnectionManage.vue";
let list = ref()
let props = defineProps(['flush'])
let infoDbList = ref()

watch(props, (newFlush)=>{
  connectionList()
})

// 连接列表
function connectionList(){
  ConnectionList().then(res => {
    if (res.code !== 200) {
      ElNotification({
        title:res.msg,
        type: "error",
      })
    }
    list.value = res.data
  })
}
connectionList()

// 删除连接
function connectionDelete(identity) {
  ConnectionDelete(identity).then(res => {
    if (res.code !== 200) {
      ElNotification({
        title:res.msg,
        type: "error",
      })
    }
    ElNotification({
      title:res.msg,
      type: "success",
    })
    connectionList()
  })
}

// 获取基本信息
function getInfo(identity) {
  // 获取数据库列表
  DbList(identity).then(res => {
    if (res.code !== 200) {
      ElNotification({
        title:res.msg,
        type: "error",
      })
      return
    }
    infoDbList.value = res.data
  })
}
</script>

<style scoped>
.item {
  display: flex;
  width: 100%;
  justify-content: space-between;
}
.db-item {
  color: #409eff;
  background-color: #ecf5ff;
  padding: 5px 12px;
  margin-bottom: 5px;
}
</style>
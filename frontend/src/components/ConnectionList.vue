<template>
  <main>
    <div class="demo-collapse">
      <el-collapse accordion>
        <el-collapse-item v-for="item in list" :name="item.identity">
          <template #title>
            <div class="item">
              <span>
                {{ item.name }}
              </span>
              <span>
                <ConnectionManage @click.stop title="编辑" btn-type="text" :data="item" @emit-connection-list="connectionList"/>
              </span>
            </div>
          </template>
        </el-collapse-item>
      </el-collapse>
    </div>
  </main>
</template>

<script setup>
import {ref, watch} from "vue";
import {ConnectionList} from "../../wailsjs/go/main/App.js";
import {ElNotification} from "element-plus"
import ConnectionManage from "./ConnectionManage.vue";
let list = ref()
let props = defineProps(['flush'])
watch(props, (newFlush)=>{
  connectionList()
})
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
</script>

<style scoped>
.item {
  display: flex;
  width: 100%;
  justify-content: space-between;
}
</style>
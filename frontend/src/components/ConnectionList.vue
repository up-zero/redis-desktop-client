<template>
  <main>
    <div class="demo-collapse">
      <el-collapse accordion>
        <el-collapse-item v-for="item in list" name="1">
          <template #title>
            {{ item.name }}
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

</style>
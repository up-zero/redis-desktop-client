<template>
  <main v-if="keyKey !== undefined">
    <el-form :model="form">
      <el-form-item label="键">
        <el-input type="textarea" autosize v-model="form.key" disabled placeholder="" />
      </el-form-item>
      <el-row>
        <el-col :span="11">
          <el-form-item label="过期时间">
            <el-input v-model="form.ttl" placeholder="" />
          </el-form-item>
        </el-col>
        <el-col :span="2"></el-col>
        <el-col :span="11">
          <el-form-item label="数据类型">
            <el-input v-model="form.type" disabled placeholder="" />
          </el-form-item>
        </el-col>
      </el-row>
      <div v-if="form.type === 'string'">
        <el-form-item label="值">
          <el-input type="textarea" :autosize="{ minRows: 6 }" v-model="form.value" placeholder="" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="updateKey">保存</el-button>
        </el-form-item>
      </div>
      <div v-else-if="form.type === 'hash'">
        <el-table :data="form.value" border style="width: 100%">
          <el-table-column prop="key" label="Key" />
          <el-table-column prop="value" label="Value" />
          <el-table-column label="操作">
            <template #default="scope">
              <el-popconfirm title="确认删除?" @confirm="deleteHashField([scope.row.key])">
                <template #reference>
                  <el-button link type="danger">删除</el-button>
                </template>
              </el-popconfirm>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-form>
  </main>
</template>

<script setup>
import {ref, watch} from 'vue'
import {GetKeyValue, HashFieldDelete, UpdateKeyValue} from "../../wailsjs/go/main/App.js";
import {ElNotification} from "element-plus";
let props = defineProps(['keyDB', 'keyConnIdentity', 'keyKey'])
let form = ref({})

watch(()=>props.keyKey, () => {
  getTheValue()
})

function getTheValue() {
  GetKeyValue({conn_identity:props.keyConnIdentity, db: props.keyDB, key: props.keyKey}).then(res => {
    if (res.code !== 200) {
      ElNotification({
        title:res.msg,
        type: "error",
      })
      return
    }
    form.value = res.data
    form.value.key = props.keyKey
  })
}

function updateKey() {
  UpdateKeyValue({conn_identity: props.keyConnIdentity, key: props.keyKey, db: props.keyDB, value: form.value.value, ttl: form.value.ttl}).then(res => {
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
  })
}

function deleteHashField(fields) {
  HashFieldDelete({conn_identity: props.keyConnIdentity, db: props.keyDB, key: props.keyKey, field:fields}).then(res => {
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
    getTheValue()
  })
}
</script>

<style scoped>

</style>
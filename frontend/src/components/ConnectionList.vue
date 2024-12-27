<script setup>

import {ref, watch} from "vue";
import {ConnectionDelete, ConnectionList, DbList} from "../../wailsjs/go/main/App.js";
import {ElNotification} from "element-plus";
import ConnectionManage from "./ConnectionManage.vue";

let list = ref()
let infoDbList = ref()
let selectDbKey = ref()
let emits = defineEmits(['emit-select-db'])
let props = defineProps(['flush'])
watch(props, (newFlush) => {
  connectionList()
})

function getInfo(identity){
  infoDbList.value = []
  DbList(identity).then(res=>{
    if(res.code !==200){
      ElNotification({
        title: res.msg,
        type: "error",
      })
    }
    infoDbList.value = res.data
  })
}

// 获取连接列表
function connectionList(){
  ConnectionList().then((res) => {
    if(res.code !==200){
      ElNotification({
        title: res.msg,
        type: "error",
      })
    }
    list.value = res.data
  })
}
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
// 选中数据库
function selectDB(db , identity){
  selectDbKey.value = db
  // db0 db1
  // 0 1
  emits('emit-select-db', Number(db.substring(2)), identity)
}

connectionList()
</script>

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
                <ConnectionManage @click.stop title="编辑" btn-type="text" :data="item" @emit-connection-list="connectionList" />
                <el-popconfirm title="确定删除吗？" @confirm="connectionDelete(item.identity)">
                  <template #reference>
                    <el-button link type="danger" @click.stop>删除</el-button>
                  </template>
                </el-popconfirm>
              </div>
            </div>
          </template>
          <div v-for="db in infoDbList" class="my-item" @click="selectDB(db.key,item.identity)">
            <div v-if="db.key !== selectDbKey" class="my-item"> {{db.key}} ({{db.number}})</div>
            <div v-else class="my-select-item"> {{db.key}} ({{db.number}})</div>
          </div>

        </el-collapse-item>

      </el-collapse>
    </div>
  </main>
</template>

<style scoped>
.item{
  display: flex;
  justify-content: space-between;
  width: 100%;
}

</style>
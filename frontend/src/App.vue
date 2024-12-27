<script setup>
import HelloWorld from './components/HelloWorld.vue'
import ConnectionList from "./components/ConnectionList.vue";
import ConnectionManage from "./components/ConnectionManage.vue";
import {ref} from "vue";
import Keys from "./components/keys.vue";

let flushFlag = ref(true)
let keyDB = ref()
let keyConnIdentity = ref()
let keyKey = ref()

function flushConnectionList(){
  flushFlag.value = !flushFlag.value
}
function selectDB(db, connIdentity){
  keyDB.value = db
  keyConnIdentity.value = connIdentity
}
</script>

<template>
  <el-row>
    <el-col :span="5" style="height: 100vh; padding:12px ">
      <div style="margin-bottom: 12px">
        <ConnectionManage title="新建连接" btn-type="primary" @emit-connection-list="flushConnectionList" />
      </div>
      <ConnectionList @emit-select-db="selectDB" :flush="flushFlag"/>
    </el-col>
    <el-col :span="9" style="padding: 12px">
      <Keys :keyDB="keyDB" :keyConnIdentity="keyConnIdentity"/>
    </el-col>
  </el-row>


</template>

<style>

</style>

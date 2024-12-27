<script setup>
import {reactive, ref} from 'vue'
import {ConnectionList,ConnectionCreate} from '../../wailsjs/go/main/App'


const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
})
let code =ref()
let list = ref()
let createres = ref()
function greet() {
  ConnectionCreate({"addr":"127.0.0.1:8080"}).then((res) => {
    createres.value = res
  })
  ConnectionList().then((res) => {
    code.value = res.code
    list.value = res.data
  })

}

</script>

<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <input id="name" v-model="data.name" autocomplete="off" class="input" type="text"/>
      <button class="btn" @click="greet">Greet</button>
      <div>code:  {{code}}</div>
      <div>list:  {{list}}</div>
      <div>createres:  {{createres}}</div>
    </div>

  </main>
</template>

<style scoped>
.result {
  height: 20px;
  width: 100%;
  text-align: center;
  line-height: 20px;
  margin: 1.5rem auto;
  color: #cc7c7c;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  text-align: center;

  background-image: linear-gradient(to top, #7746d0 0%, #1d5170 100%);
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #1d5170 0%, #7746d0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  width: 300px;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgb(227, 205, 205);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>

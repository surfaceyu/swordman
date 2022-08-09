<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://vuejs.org/api/sfc-script-setup.html#script-setup
</script>

<template>
  <router-view></router-view>
</template>

<style>
p {
  font-size: 21px;
  line-height: 1.6em;
  color: #000;
  font-family: '楷体';
}

a:link {
  color: #4e6ef2;
  font-weight: normal;
  text-decoration: none;
}

@keyframes dh {
  0% {
    background-position: 0 0;
  }

  100% {
    background-position: -100% 0;
  }
}

a .a_link_1 {
  color: blue
}

a .a_link_2 {
  color: #20B2AA
}

a .a_link_3 {
  color: #8A2BE2
}

a .a_link_4 {
  color: #FF8C00;
  animation: dh 4s infinite linear
}

a .a_link_5 {
  color: #DAA520;
  animation: dh 4s infinite linear
}

a .a_link_6 {
  background: -webkit-linear-gradient(left, #3498db, #f47920 10%, #d71345 20%, #f7acbc 30%, #ffd400 40%, #3498db 50%, #f47920 60%, #d71345 70%, #f7acbc 80%, #ffd400 90%, #3498db);
  background-clip: text;
  animation: dh 4s infinite linear;
  color: transparent;
  background-size: 200% 100%;
}

.txt_1 {
  color: blue
}

.txt_2 {
  color: #20B2AA
}

.txt_3 {
  color: #8A2BE2
}

.txt_4 {
  color: #FF8C00
}

.txt_5 {
  color: #DAA520;
}

.a_btn {
  outline: none;
  border: none;
  color: #4e6ef2;
  font-weight: normal;
  text-decoration: none;
  font-size: 21px;
  line-height: 1.6em;
  font-family: '楷体';
  background: rgba(0, 0, 0, 0);
  padding: 0;
}

.txt_6 {
  background: linear-gradient(45deg, #DC143C, #FF1493, #00BFFF, #1E90FF);
  background-clip: text;
  color: transparent;
}
</style>
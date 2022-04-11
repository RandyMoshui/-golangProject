<template>
<div id="container">

  <el-drawer
  title="存活探测详细信息"
  :visible.sync="sniff_drawer"
  :direction="ttb"
  :before-close="handleClose"
  size="100%"
  :modal=false
  modal-append-to-body=false
  show-close=true
  style="opacity:0.9;"
  >
  <SniffDrawer>

  </SniffDrawer>
</el-drawer>
    <!-- 扫描 -->
    <SniffCard v-on:open_drawer="sniff_drawer=true;">

    </SniffCard>

    <!-- 渗透测试 -->
    <TestCard></TestCard>
    <!-- 安全情报 -->
    <FeelCard></FeelCard>

    

</div>
</template>

<script>

import SniffCard from "./ClipCard/SniffCard.vue";
import TestCard from "./ClipCard/TestCard.vue";
import FeelCard from './ClipCard/FeelCard.vue';

import SniffDrawer from "./drawer/SniffDrawer.vue"

export default {
  name: 'IndexContainer',
  components: {
    SniffCard,
    TestCard,
    FeelCard,
    SniffDrawer,
  },
  data:function(){
    return {

      sniff_drawer: false,
    }
  },
  methods: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

#container {
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content:space-around;
  flex-wrap: wrap;
  align-items: center;
}

</style>

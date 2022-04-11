<template>
   <!-- 态势感知  -->
    <div class="card" id="feel">
        <div id="front">
            <h2>安全情报</h2>
        </div>

        <div id="back">
            2{{value}}
            <el-switch
            v-model="value"
            active-color="#13ce66"
            inactive-color="#ff4949">
            </el-switch>
        </div>
    </div>
</template>

<script>

export default {
  name: 'FeelCard',
  components: {
    
  },
  data:function(){
    return {
      value: true,
      drawer: false,
    }
  },
  methods: {
      
  }
}

</script>


<style scoped>
@import url(../../../public/card.css);

#back{
    justify-content:center;
    align-items:center;
}
</style>
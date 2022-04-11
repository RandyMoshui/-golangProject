<template>
  <!-- 渗透测试  -->
    <div class="card" id="test">
        <div id="front">
            <h2>渗透测试</h2>
        </div>

        <div id="back" @click="cancel_hvoer();">
            <div class="btn_container">
                <div class="btn" style="font-weight: 800;">端口扫描</div>
                <div class="btn" style="font-weight: 800;">漏洞测试</div>
                <div class="btn" style="font-weight: 300;">more</div>
            </div>

            <div class="content">
                <div id="port_scan">
                    <span>IP地址(段):</span>
                     <el-tooltip class="item" effect="dark" content="输入示例:192.168.2.1-192.168.2.1" placement="top">
                     <el-input
                        placeholder="请输入内容"
                        v-model="ip_address"
                        size="small"
                        clearable
                     >
                    </el-input>
                    </el-tooltip>

                    <el-dropdown trigger="click">
                        <span class="el-dropdown-link">
                            选择扫描端口
                            <i class="el-icon-arrow-down el-icon--right"></i>
                        </span>
                        <el-dropdown-menu slot="dropdown">
                            <el-dropdown-item icon="el-icon-plus">常用端口</el-dropdown-item>
                            <el-dropdown-item icon="el-icon-circle-plus">常用端口+随机廥</el-dropdown-item>
                            <el-dropdown-item icon="el-icon-circle-plus-outline">所有端口</el-dropdown-item>
                        </el-dropdown-menu>
                    </el-dropdown>
                </div>

                <div id="leak_test">

                </div>
            </div>
            
        </div>
    </div>
</template>

<script>

export default {
  name: 'TestCard',
  components: {
    
  },
  data:function(){
    return {
        ip_address: '',
    }
  },
  methods: {
      cancel_hvoer(){
          document.querySelector('#back').classList.add('disappear');
      }
  }
}

</script>


<style scoped>
@import url(../../../public/card.css);

#back{
    width: 100%;
    flex-direction: column;
    justify-content:flex-start;
    align-items: flex-start;
}

#back .btn_container{
    width: 100%;
    display:flex;
    flex-direction:row;
    justify-content:center;
    padding: 5px;
}

#back .btn_container .btn{
    margin: 5px;
    border-bottom: solid rgba(16,174,194,0.2);
    color: rgba(0,0,0,0.8)
}

#back .btn_container .btn:hover{
    color: rgb(18, 110, 130, 0.2);
    cursor: default;
}

#back .content{
    width: 85%;
    height: 100%;

    display: flex;
    flex-direction: column;

    padding: 25px;
}

.disappear{
    display:none;
    
}



</style>
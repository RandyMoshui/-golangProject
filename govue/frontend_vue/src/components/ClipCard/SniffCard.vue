<template>
    <!-- 扫描（存活与端口） -->
  <div class="card" id="sniff">
        <div class="front1" :id="front_img_path">
            <img src="" alt="">
            <h2>扫描</h2>
            
        </div>

        <div class="back1">
            <!-- 提交前的输入框 -->
            <div>
            <div style="width:0.8;display:flex;flex-direction:column;">
                <div id="Ip_Input">
                    <span style="margin-bottom:10px;">IP地址：</span>
                <el-tooltip content="输入示例:192.168.2.1-192.168.2.1" placement="top">
                <el-input
                    placeholder="请输入IP地址(段)"
                    v-model="ip_address"
                    @change="verify_input();"
                    size="small"
                    clearable
                    style="opacity:0.2;"
                    >
                </el-input>
                </el-tooltip>
                </div>
                
            </div>

            <!-- 提交后的进度条显示 -->
            <div id="show_info" class="disappear"
            style="text-align:center;">
                <span>扫描IP地址段为：</span>
                <span>{{ip_address}}</span>
                <br>
                <el-progress type="circle" :percentage="100" status="success"></el-progress>
            </div>

            
            <!-- 是否开启UDP穿透 -->
            <div id="UDP_Choose">
                <span >开启UDP</span>
                <el-switch
                v-model="value"
                active-color="#13ce66"
                inactive-color="#ff4949"
                style="margin: 15px;"
                >
                </el-switch>
            </div>
            

            <!-- 开始、结束按钮 -->
            <div id="button" style="display:flex;flex-direction:column;margin-top:15px;">
                <el-button
                @click="testprogress();" 
                :type.sync="click_button_type"
                style="align-self: center;">
                {{click_button_data}}
                </el-button>
                <!-- 详细信息按钮 -->
                <!-- disappear -->
                <div id="detail" class="" style="text-align:center;">
                    <el-button
                        class="detail"
                        type="primary"
                        v-on:click="$emit('open_drawer');">
                        详细信息
                    </el-button>
                </div>
                
            </div>
            

            </div>
            

        </div>
    </div>
</template>

<script>
import nprogress from 'nprogress';

export default {
  name: 'SniffCard',
  components: {
    
  },
  data:function(){
    return {
      value: true,
      drawer: false,
      ip_address: '',
      click_button_type: 'primary',
      click_button_data: '开始',
      begin: false,
    }
  },
  methods: {
    //   开始终止方法
      testprogress(){
          // 提交开始
          if(this.begin === false){
            var patt = /^(?=(\b|\D))(((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))\.){3}((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))(?=(\b|\D))-(?=(\b|\D))(((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))\.){3}((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))(?=(\b|\D))$/;
            // 验证IP地址格式
            if (patt.test(this.ip_address) === false){
                this.$prompt('请重新输入IP地址', '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    inputPattern: /^(?=(\b|\D))(((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))\.){3}((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))(?=(\b|\D))-(?=(\b|\D))(((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))\.){3}((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))(?=(\b|\D))$/,
                    inputErrorMessage: 'IP地址格式不正确'
                    }).then(({ value }) => {
                        this.$message({
                            type: 'success',
                            message: '待扫描IP段为: ' + value
                            });
                            this.ip_address = value;
                            }).catch(() => {
                                this.$message({
                                    type: 'info',
                                    message: '取消提交'
                                    });
                                    this.ip_address = '';       
                                    });
            }
            else if(patt.test(this.ip_address) === true){
                //   恢复反面效果
                document.querySelector('.front1_disappear').classList.add("front1");
                document.querySelector('.back1_persist').classList.add("back1");
                //   显示进度条，并隐藏输入框和滑动选择框、详细信息按钮
                document.querySelector('#Ip_Input').classList.add("disappear");
                document.querySelector('#UDP_Choose').classList.add("disappear");
                document.querySelector('#show_info').classList.remove("disappear");
                document.querySelector('#detail').classList.remove("disappear");

                this.begin = true;
                nprogress.configure({ parent: '.front1', showSpinner: false});
                nprogress.start();
                this.click_button_type = "danger";
                this.click_button_data = '结束';
            }       
          }
          // 提交终止
          else if(this.begin === true){
              this.begin = false;
              // 切换按钮（结束-》开始）
              this.click_button_type = "primary";
              this.click_button_data = '开始';
            nprogress.done(true); // 关闭正面进度条
            //   消除进度条，并显示输入框和滑动选择框、隐藏详细信息按钮
            document.querySelector('#Ip_Input').classList.remove("disappear");
            document.querySelector('#UDP_Choose').classList.remove("disappear");
            document.querySelector('#show_info').classList.add("disappear");
            document.querySelector('#detail').classList.add("disappear");
          }
          
      },
        // 验证输入方法
      verify_input(){
          if(document.querySelector('.front1')){
              document.querySelector('.front1').classList.add("front1_disappear");
              document.querySelector('.front1').classList.toggle("front1");
          }

          if(document.querySelector('.back1')){
              document.querySelector('.back1').classList.add("back1_persist");
              document.querySelector('.back1').classList.toggle("back1");
          }
        
        
        //   alert(this.ip_address);
        var patt = /^(?=(\b|\D))(((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))\.){3}((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))(?=(\b|\D))-(?=(\b|\D))(((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))\.){3}((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))(?=(\b|\D))$/;
        if (patt.test(this.ip_address) === false){
            this.$prompt('请重新输入IP地址', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          inputPattern: /^(?=(\b|\D))(((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))\.){3}((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))(?=(\b|\D))-(?=(\b|\D))(((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))\.){3}((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))(?=(\b|\D))$/,
          inputErrorMessage: 'IP地址格式不正确'
        }).then(({ value }) => {
          this.$message({
            type: 'success',
            message: '待扫描IP段为: ' + value
          });
          this.ip_address = value;
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '取消输入'
          });
          this.ip_address = '';       
        });
        }
        
      },

  }
}

</script>


<style scoped>
.card{
    display:flex;
    flex-direction:column;
    justify-content:center;
    align-items:center;

    position: relative;

    width: 360px;
    height: 500px;
    margin: 15px;
}

.front1_disappear{
    display: none;
}

.back1_persist{
    width: 100%;
    height: 100%;

    position: absolute;

    border-radius: 25px;
    /* background: #2f363e;
    box-shadow: 25px 25px 35px rgba(0, 0, 0, 0.25),
    10px 10px 30px rgba(0, 0, 0, 0.25),
    inset 5px 5px 5px rgba(0, 0, 0, 0.5),
    inset 5px 5px 15px rgba(255,255,255,0.2),
    inset -5px -5px 15px rgba(0, 0, 0, 0.75); */
    background: linear-gradient(
        to right bottom,
        rgba(255, 255, 255, 0.6),
        rgba(255, 255, 255, 0.1)
    );
    backdrop-filter: blur(8px);
    box-shadow: 10px 10px 20px rgba(0, 0, 0, 0.2);
    border-top: 1px solid rgba(255, 255, 255, 0.8);
    border-left: 1px solid rgba(255, 255, 255, 0.8);
    
    display: flex;
    justify-content:center;
    align-items:center;
    flex-direction:column;

    background-size: cover;
    backface-visibility: hidden;

    /* 设置动画变化过程 */
    transition: 0.8s;
}

.front1, .back1{
    width: 100%;
    height: 100%;

    position: absolute;

    border-radius: 25px;
    /* background: #2f363e;
    box-shadow: 25px 25px 35px rgba(0, 0, 0, 0.25),
    10px 10px 30px rgba(0, 0, 0, 0.25),
    inset 5px 5px 5px rgba(0, 0, 0, 0.5),
    inset 5px 5px 15px rgba(255,255,255,0.2),
    inset -5px -5px 15px rgba(0, 0, 0, 0.75); */
    background: linear-gradient(
        to right bottom,
        rgba(255, 255, 255, 0.6),
        rgba(255, 255, 255, 0.1)
    );
    backdrop-filter: blur(8px);
    box-shadow: 10px 10px 20px rgba(0, 0, 0, 0.2);
    border-top: 1px solid rgba(255, 255, 255, 0.8);
    border-left: 1px solid rgba(255, 255, 255, 0.8);
    
    display: flex;
    justify-content:center;
    align-items:center;
    flex-direction:column;

    background-size: cover;
    backface-visibility: hidden;

    /* 设置动画变化过程 */
    transition: 0.8s;
}

.front1 {
    transform: perspective(800px) rotateY(0deg);
    user-select: none;
}

.back1 {
    transform: perspective(800px) rotateY(180deg);
}

.card:hover .front1{
    transform: perspective(800px) rotateY(-180deg);
}

.card:hover .back1{
    transform: perspective(800px) rotateY(0deg);
}

.card .detail {
    margin-top:15px;
    align-self: center;
}

.disappear{
    display: none;
}
</style>
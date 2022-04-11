<template>
    <div class="container">
        <div class="right">
                <div class="rt">
                    <el-progress type="circle" :percentage="100" status="success">

                    </el-progress>
                <br>
                剩余未扫主机数：
                </div>

                <div class="rb">
                    <el-progress type="circle" :percentage="100" status="success">
                    </el-progress>
                    <br>
                    已扫描存活主机数：
                </div>
                

            
        </div>

        <div class="left">
            <el-steps :active="2" align-center style="width:80%;margin:15px;">
                <el-step title="步骤1" description="程序准备"></el-step>
                <el-step title="步骤2" description="Ping"></el-step>
                <el-step title="步骤3" description="存活扫描"></el-step>
                <el-step title="步骤4" description="UDP扫描"></el-step>
            </el-steps>

            <div>
                <el-table
            :data="tableData"
            style="width: 100%;margin-top:15px;"
            height="250">
                <el-table-column
                fixed
                prop="date"
                label="IP地址"
                width="150">
                </el-table-column>
                <el-table-column
                prop="name"
                label="ARP"
                width="120">
                </el-table-column>
                <el-table-column
                prop="province"
                label="Ping"
                width="120">
                </el-table-column>
                <el-table-column
                prop="city"
                label="UDP"
                width="120">
                </el-table-column>
                <el-table-column
                fixed
                prop="address"
                label="存活判断"
                width="120">
                </el-table-column>
            </el-table>
            </div>
            
        </div>
        
    </div>
</template>

<script>

export default{
    name: 'SniffCard',
}

</script>

<style scoped>
.container {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;

    border-radius: 20px;
    /* 设置背景渐变 */
    background: linear-gradient(
        to right bottom,
        rgba(255, 255, 255, 0.6),
        rgba(255, 255, 255, 0.1)
    );
    /* 设置背景模糊 */
    backdrop-filter: blur(8px);
    /* 设置阴影 */
    box-shadow: 10px 10px 20px rgba(0, 0, 0, 0.2);

    /* 渐变边框（感觉可有可无） */
    /* border-top: 1px solid rgba(255, 255, 255, 0.8);
    border-left: 1px solid rgba(255, 255, 255, 0.8); */

    color: black;
}

.right {
    min-height: 80%;
    flex: 1;

    display: flex;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;

    border-radius: 25px;
}

.left {
    min-height: 80%;
    flex: 2;

    display:flex;
    flex-direction:column;
    justify-content: space;
    align-items: center;
    overflow: hidden;
    margin-right: 5px;

    /* border: solid; */
    border-radius: 25px;

    box-shadow: 0px 0px 10px 5px rgb(7, 7, 7,0.1);
    
}

.rt, .rb{
    display: flex;
    flex-direction:column;
    justify-content:center;
    align-items: center;

    padding: 20px;
    border-radius: 25px;

    box-shadow: 0px 0px 10px 5px rgb(7, 7, 7,0.1);
}

</style>

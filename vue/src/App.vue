<template>
    <div id="app">
        <div class="upload-box">
            <el-upload class="upload-demo"
                       drag
                       action="/api/upload"
                       :show-file-list="false"
                       :on-success="handleSuccess"
                       multiple>
                <i class="el-icon-upload"></i>
                <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
            </el-upload>
            <div class="copy">
                <el-input placeholder="下载地址"
                          disabled
                          v-model="currentUrl">
                </el-input>
                <el-button style="margin-left:5px"
                           type="success"
                           :disabled="currentUrl == ''"
                           @click="handleDownload(currentUrl)">
                    下载
                </el-button>
            </div>
        </div>
        <h5 class="his">历史记录:显示最近100条</h5>
        <el-table border
                  :data="list">
            <el-table-column label="序号"
                             type="index"
                             align="center"
                             width="60px" />
            <el-table-column align="center"
                             label="zip文件">
                <template slot-scope="scope">
                    <div class="img-box">
                        <img src="./zip.png">
                        <p>{{scope.row.file}}</p>
                    </div>
                </template>
            </el-table-column>
            <el-table-column align="center"
                             label="操作"
                             width="200px">
                <template slot-scope="scope">
                    <el-button @click="handleDownload(scope.row.file)"
                               size="mini"
                               type="success">下载</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
import request from 'axios'
export default {
    name: 'App',
    data() {
        return {
            currentUrl: '',
            list: []
        }
    },
    async created() {
        const { data } = await request.get('/api/history')
        if (data.code == 1) {
            data.data.forEach(v => {
                this.list.push({ file: v })
            })
        } else {
            this.$message.error(data.msg)
        }
    },
    methods: {
        // 文件上传成功
        handleSuccess(res) {
            if (res.code == 1) {
                this.currentUrl = res.data
                this.list.unshift({ file: res.data })
            } else {
                this.$message.error(res.msg)
            }
        },
        handleDownload(file) {
            console.log(file)
            window.open("www.baidu.com")
        }
    }
}
</script>

<style>
body {
    padding: 100px;
}
.upload-box {
    display: flex;
    align-items: center;
}
.copy {
    flex: 1;
    padding: 20px;
    display: flex;
}
.his {
    color: #ccc;
}
.el-icon-upload {
    margin: 0 !important;
}
.el-upload-dragger {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}
.img-box {
    margin: 0 auto;
    display: flex;
    justify-content: center;
    align-items: center;
}
.img-box img {
    width: 40px;
    height: 40px;
}
.img-box p {
    margin-left: 20px;
}
</style>

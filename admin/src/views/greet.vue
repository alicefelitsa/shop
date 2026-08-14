<template>
  <div style="width: 100%;">
    <el-main>
      <el-card class="box-card" shadow="always">
        <div slot="header">
          <el-button type="primary" plain @click="dialogVisible = true">添加打招呼</el-button>
        </div>
        <el-table :data="tableData" style="width: 100%;" height="calc(100vh - 150px)" :border="true"
                  v-loading="visitorLoading">
          <el-table-column prop="content" label="内容" align="center">
            <template v-slot="{row}">
              <div v-if="row.types==='text'" class="content" v-html="row.content"></div>
              <div v-else-if="row.types==='image'" class="content">
                <el-image
                    style="width: 200px;"
                    :src="file_url+row.content"
                    :preview-src-list="[file_url+row.content]">
                </el-image>
              </div>
              <div v-else-if="row.types==='video'" class="content">
                <video :key="row.content" controls style="width: 100%;height: 300px;">
                  <source :src="file_url+row.content" type="video/mp4">
                </video>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="sort" label="顺序" align="center" width="220px">
            <template v-slot="{row}">
              <el-input-number v-model="row.sort" @change="handleSort(row)" :min="0" :max="100"></el-input-number>
            </template>
          </el-table-column>
          <el-table-column prop="status" label="状态" align="center" width="100px">
            <template v-slot="{row}">
              <el-switch
                  v-model="row.status"
                  :active-value="0"
                  :inactive-value="1"
                  @change="handleStatus(row)"
              >
              </el-switch>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" width="180px">
            <template v-slot="{row}">
              <el-button
                  size="mini"
                  @click="edit(row)">编辑
              </el-button>
              <el-button
                  size="mini"
                  type="danger"
                  @click="del(row)">删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-card>
    </el-main>
    <!--编辑打招呼对话框-->
    <el-dialog :title="editData?'修改打招呼':'添加打招呼'" :visible.sync="dialogVisible" width="50%"
               :close-on-click-modal="false" :before-close="handleClose" top="8vh">
      <el-tabs type="border-card" v-model="activeName" @tab-click="handleClick">
        <el-tab-pane label="文本" name="text" style="margin-bottom: 65px">
          <quill-editor
              style="height: 300px;"
              v-model="content"
              :options="editorOptions"
              @change="onEditorChange"
          />
        </el-tab-pane>
        <el-tab-pane label="图片" name="image">
          <el-upload
              class="avatar-uploader"
              :action="file_url+'/uploading/api/uploadImages'"
              :show-file-list="false"
              accept="image/*"
              :on-success="handleAvatarSuccess"
              :before-upload="beforeAvatarUpload"
          >
            <img v-if="uploadImage" :src="file_url+uploadImage" class="avatar_greet" alt="">
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-tab-pane>
        <el-tab-pane label="视频" name="video">
          <el-upload
              class="avatar-uploader"
              :action="file_url+'/uploading/api/uploadFiles'"
              :show-file-list="false"
              accept="video/mp4"
              :on-success="handleAvatarSuccessForVideo"
              :before-upload="beforeAvatarUploadForVideo"
          >
            <video v-if="uploadVideo" :key="uploadVideo" controls style="width: 100%;height: 308px;">
              <source :src="file_url+uploadVideo" type="video/mp4">
            </video>
            <el-button type="primary" plain>上传视频</el-button>
          </el-upload>
        </el-tab-pane>
      </el-tabs>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleClose">取 消</el-button>
        <el-button type="primary" @click="save">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import {delGreet, getGreetList, saveGreet, saveGreetSort, saveGreetStatus} from "@/api/greet";
import Quill from 'quill'
import QuillEmoji from 'quill-emoji'
import 'quill-emoji/dist/quill-emoji.css'

Quill.register({
  'modules/emoji-toolbar': QuillEmoji.ToolbarEmoji,
  'modules/emoji-shortname': QuillEmoji.ShortNameEmoji
})
export default {
  data() {
    return {
      loading: null,
      tableHeight: '80vh',
      uploadImage: '',
      uploadVideo: '',
      editData: null,
      file_url: '',
      tableData: [],
      visitorLoading: false,
      dialogVisible: false,
      activeName: 'text',
      content: '',
      editorOptions: {
        theme: 'snow', // 主题：snow（默认）、bubble
        placeholder: '请输入内容...',
        modules: {
          'emoji-toolbar': true,    // 显示表情选择面板
          'emoji-shortname': true,  // 支持短名称（如 :smile:）
          toolbar: [
            ['bold', 'italic', 'underline', 'strike'],        // 加粗、斜体、下划线、删除线
            //['blockquote', 'code-block'],                      // 引用、代码块
            //[{'header': 1}, {'header': 2}],                // 标题1、标题2
            //[{'list': 'ordered'}, {'list': 'bullet'}],     // 有序列表、无序列表
            //[{'script': 'sub'}, {'script': 'super'}],      // 下标、上标
            //[{'indent': '-1'}, {'indent': '+1'}],          // 缩进
            [{'direction': 'rtl'}],                         // 文字方向
            [{'size': ['small', false, 'large', 'huge']}],  // 字体大小
            //[{'header': [1, 2, 3, 4, 5, 6, false]}],        // 标题级别
            [{'font': []}],                                 // 字体
            [{'color': []}, {'background': []}],          // 字体颜色、背景色
            [{'align': []}],                                // 对齐方式
            //['clean'],                                         // 清除格式
            //['link', 'image', 'video']                        // 链接、图片、视频
            ['emoji']
          ]
        }
      }
    }
  },
  mounted() {
    this.getGreet()
  },
  methods: {
    //获取主动打招呼
    getGreet() {
      this.visitorLoading = true;
      getGreetList(localStorage.getItem("ucode")).then((data) => {
        //console.log(data);
        setTimeout(() => {
          this.visitorLoading = false;
          this.tableData = data.data;
          this.file_url = data.configData.file_url;
        }, 200)
      }).catch((e) => {
        this.$message.error(e.message);
      });
    },
    //设置排序
    handleSort(row) {
      saveGreetSort({
        ucode: localStorage.getItem("ucode"),
        sort: row.sort,
        id: row.id,
      }).then((msg) => {
        this.$message.success(msg)
        //this.getGreet()
      }).catch((e) => {
        this.$message.error(e.message);
      });
      //console.log(row.id, row.sort);
    },
    //设置开启状态
    handleStatus(row) {
      //console.log(row.id, row.status);
      saveGreetStatus({
        ucode: localStorage.getItem("ucode"),
        status: row.status,
        id: row.id,
      }).then((msg) => {
        this.$message.success(msg)
        //this.getGreet()
      }).catch((e) => {
        row.status = !row.status ? 1 : 0;
        this.$message.error(e.message);
      });
    },
    edit(row) {
      this.editData = row
      this.dialogVisible = true;
      if (row.types === 'text') {
        this.content = row.content;
      } else if (row.types === 'image') {
        this.uploadImage = row.content;
      } else if (row.types === 'video') {
        this.uploadVideo = row.content;
      }
      this.activeName = row.types;
      //console.log(row)
    },
    save() {
      let saveStatus
      if (this.editData == null) {
        saveStatus = 'add'
      } else {
        saveStatus = 'edit'
      }
      saveGreet(localStorage.getItem("ucode"), {
        saveStatus: saveStatus,
        types: this.activeName,
        content: this.content,
        uploadImage: this.uploadImage,
        uploadVideo: this.uploadVideo,
        id: this.editData?.id,
      }).then((msg) => {
        this.handleClose()
        this.$message.success(msg)
        this.getGreet()
      }).catch((e) => {
        this.$message.error(e.message);
      });
    },
    del(row) {
      this.$confirm('确认删除？').then(_ => {
        delGreet({
          ucode: localStorage.getItem("ucode"),
          id: row.id,
        }).then((msg) => {
          this.$message.success(msg)
          this.getGreet()
        }).catch((e) => {
          this.$message.error(e.message);
        });
      }).catch(_ => {
      });
    },
    onEditorChange({quill, html, text}) {
      // console.log('编辑器内容发生变化')
      // console.log('HTML:', html)
      // console.log('纯文本:', text)
      this.content = html
    },
    //关闭编辑对话框
    handleClose() {
      this.dialogVisible = false
      this.content = ''
      this.editData = null
      this.uploadImage = ''
      this.uploadVideo = ''
      this.activeName = 'text'
    },
    handleClick(tab, event) {
      console.log(this.activeName);
      //console.log(tab.name);
    },
    //图片上传成功
    handleAvatarSuccess(response, file, fileList) {
      if (response.code === 0) {
        this.uploadImage = response.filePath
      } else {
        this.$message.error(response.message)
      }
    },
    //视频上传成功
    handleAvatarSuccessForVideo(response, file, fileList) {
      if (response.code === 0) {
        this.uploadVideo = response.filePath
        console.log("上传视频：", this.uploadVideo)
      } else {
        this.$message.error(response.message)
      }
      this.$nextTick(() => {
        this.loading.close()
      })
    },
    //上传图片前钩子
    beforeAvatarUpload(file) {
      const isImage = file.type.startsWith('image/')
      const isLt2M = file.size / 1024 / 1024 < 3
      if (!isImage) {
        this.$message.error('只能上传图片文件!')
        return false
      }
      if (!isLt2M) {
        this.$message.error('图片大小不能超过3MB!')
        return false
      }
      return true
    },
    //上传视频前钩子
    beforeAvatarUploadForVideo(file) {
      const isMp4 = file.type === 'video/mp4'
      const isLt50M = file.size / 1024 / 1024 < 20 // 限制 50MB
      if (!isMp4) {
        this.$message.error('只能上传MP4视频!')
        return false
      }
      if (!isLt50M) {
        this.$message.error('视频大小不能超过20MB!')
        return false
      }
      this.loading = this.$loading({
        lock: true,
        text: 'Loading...视频上传中',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      });
      return true
    },
  }
}
</script>


<style>
.box-card {
  width: 98%;
  margin: 10px;
}

.el-main {
  padding: 0;
}

.content {
  margin: 0 auto;
  text-align: left;
  width: 80%;
  background-color: #fff;
  border: 1px solid #eee;
  padding: 10px;
  border-radius: 5px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, .1);
  -webkit-box-shadow: 0 2px 12px 0 rgba(0, 0, 0, .1);
}

.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

.avatar-uploader .el-upload:hover {
  border-color: #409EFF;
}

.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}

.avatar_greet {
  height: 351px !important;
  display: block !important;
}

.avatar-uploader .el-upload {
  padding: 4px;
}
</style>
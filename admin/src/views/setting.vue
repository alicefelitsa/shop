<template>
  <div style="width:100%;">
    <el-main>
      <el-card class="box-card" shadow="always">
        <div slot="header">
          <span>系统设置</span>
        </div>
        <el-form ref="form" :model="form" label-width="80px" style=" height: calc(100vh - 131px)" v-loading="visitorLoading">
          <el-row>
            <el-col :span="12">
              <el-form-item label="账号/卡号">
                <el-tag class="talName">{{ form.number }}</el-tag>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <!--              <el-form-item label="活动时间">
                              <el-col :span="11">
                                <el-date-picker type="date" placeholder="选择日期" v-model="form.date1" style="width: 100%;"></el-date-picker>
                              </el-col>
                              <el-col :span="1" style="text-align: center">-</el-col>
                              <el-col :span="11">
                                <el-time-picker placeholder="选择时间" v-model="form.date2" style="width: 100%;"></el-time-picker>
                              </el-col>
                            </el-form-item>-->
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="客服编号">
                <el-tag class="talName">{{ form.ucode }}</el-tag>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <!--              <el-form-item label="活动时间">
                              <el-col :span="11">
                                <el-date-picker type="date" placeholder="选择日期" v-model="form.date1" style="width: 100%;"></el-date-picker>
                              </el-col>
                              <el-col :span="1" style="text-align: center">-</el-col>
                              <el-col :span="11">
                                <el-time-picker placeholder="选择时间" v-model="form.date2" style="width: 100%;"></el-time-picker>
                              </el-col>
                            </el-form-item>-->
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="到期时间">
                <el-tag class="talName">{{ form.endtime }}</el-tag>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <!--              <el-form-item label="活动区域">
                              <el-select v-model="form.region" placeholder="请选择活动区域">
                                <el-option label="区域一" value="shanghai"></el-option>
                                <el-option label="区域二" value="beijing"></el-option>
                              </el-select>
                            </el-form-item>-->
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="客服昵称">
                <el-input v-model="form.nickname" style="width: 300px"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <!--              <el-form-item label="活动形式">
                              <el-input type="textarea" v-model="form.desc"></el-input>
                            </el-form-item>-->
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="客服头像">
                <el-upload
                    class="avatar-uploader"
                    :action="file_url+'/uploading/api/uploadImages'"
                    :show-file-list="false"
                    :on-success="handleAvatarSuccess"
                    :before-upload="beforeAvatarUpload"
                >
                  <img v-if="form.avatar" :src="file_url+form.avatar" class="avatar_user" alt="">
                  <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                </el-upload>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <!--              <el-form-item label="活动形式">
                              <el-input type="textarea" v-model="form.desc"></el-input>
                            </el-form-item>-->
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="访问状态">
                <el-radio-group v-model="form.visitstatus">
                  <el-radio :label="0">开启访问</el-radio>
                  <el-radio :label="1">停止访问</el-radio>
                </el-radio-group>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <!--              <el-form-item label="即时配送">
                              <el-switch v-model="form.delivery"></el-switch>
                            </el-form-item>-->
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="24">
              <el-form-item>
                <el-button type="primary" @click="onSubmit">提交保存</el-button>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
      </el-card>
    </el-main>
  </div>
</template>

<script>

import {getUserConfigData, saveUserConfig} from "@/api/setting";
import {EventBus} from "@/utils/event-bus";

export default {
  data() {
    return {
      form: {
        number: 'xxxxxxxxxxxx',
        ucode: 'xxxxxxxxxxxx',
        nickname: '在线客服',
        endtime: 'xxxxxxxxxxxx',
        avatar: '',
        visitstatus: 0
      },
      file_url: '',
      visitorLoading: false,
    }
  },
  mounted() {
    this.getUserConfig()
  },
  methods: {
    getUserConfig() {
      this.visitorLoading = true;
      getUserConfigData().then((data) => {
        //console.log(data);
        setTimeout(() => {
          this.visitorLoading = false;
          this.form = data.userData;
          this.file_url = data.configData.file_url;
        }, 200)
      }).catch((e) => {
        this.$message.error(e.message);
      });
    },
    //上传图片前钩子
    beforeAvatarUpload(file) {
      const size = file.size / 1024 / 1024 < 1;
      if (!size) {
        this.$message.error('上传图片大小不能超过 1MB');
        return false
      }
    },
    //图片上传成功
    handleAvatarSuccess(response, file, fileList) {
      if (response.code === 0) {
        this.form.avatar = response.filePath
      } else {
        this.$message.error(response.message)
      }
    },
    onSubmit() {
      saveUserConfig(this.form).then((msg) => {
        localStorage.setItem("nickname", this.form.nickname);
        localStorage.setItem("avatar", this.file_url + this.form.avatar);
        this.$message.success(msg)
        EventBus.$emit('sidebar', true)
        this.getUserConfig()
      }).catch((e) => {
        this.$message.error(e.message);
      });
    }
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

.avatar_user {
  height: 178px !important;
  width: 178px !important;
  display: block !important;
}

.talName {
  font-size: 14px;
}
</style>
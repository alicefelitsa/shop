<template>
  <div>
    <el-dialog :title="form.id?'修改产品':'添加产品'" :visible.sync="dialogVisible" width="45%"
               class="responsive-dialog" :close-on-click-modal="false" :before-close="handleClose" top="5vh">
      <el-form ref="form" :model="form" label-width="80px">
        <el-row>
          <el-col :span="24">
            <el-form-item label="名称">
              <el-input v-model="form.name" placeholder="请输入产品名称"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="分类">
              <el-select v-model="form.category" placeholder="请选择分类" style="width: 100%;">
                <el-option v-for="item in categoryList" :key="item.id" :label="item.name"
                           :value="item.name"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="价格">
              <el-input v-model="form.price" placeholder="请输入价格"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="级别">
              <el-input v-model="form.level" placeholder="请输入级别"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="纯度">
              <el-input v-model="form.purity" placeholder="请输入纯度"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="图片" class="album-item">
              <div class="album-wrap">
                <el-upload
                    class="album-uploader"
                    :action="uploadAction"
                    :headers="uploadHeaders"
                    name="file"
                    accept="image/*"
                    :show-file-list="false"
                    :before-upload="beforeUpload"
                    :on-success="onUploadSuccess"
                    :on-error="onUploadError">
                  <img v-if="form.album" :src="albumUrl" class="album-preview" alt="产品图片">
                  <div v-else class="album-placeholder">
                    <i class="el-icon-plus"></i>
                    <span>上传图片</span>
                  </div>
                </el-upload>
                <div class="album-side">
                  <div class="album-tip">支持 png/jpg/gif/jpeg，不超过 2M；点击图片可重新上传</div>
                  <el-button v-if="form.album" size="mini" icon="el-icon-delete" @click="form.album = ''">删除图片
                  </el-button>
                </div>
              </div>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="24">
            <el-form-item label="简介">
              <el-input v-model="form.Introduction" type="textarea" :rows="3" placeholder="请输入简介"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="详情">
              <el-input v-model="form.details" type="textarea" :rows="5" placeholder="请输入详情，换行分段"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="handleClose" size="small">取 消</el-button>
        <el-button type="primary" @click="save" size="small">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import {AddProduct, SaveProduct} from "@/api/productAdmin";

// 添加或修改
export default {
  props: {
    //弹窗是否打开
    visible: {
      type: Boolean,
      default: false
    },
    //修改回显的数据
    editData: {
      type: Object,
      default: () => ({})
    },
    //分类列表
    categoryList: {
      type: Array,
      default: () => ([])
    }
  },
  data() {
    return {
      form: {
        id: '',
        name: '',
        category: '',
        price: '',
        level: '',
        purity: '',
        album: '',
        Introduction: '',
        details: '',
      }
    }
  },
  computed: {
    dialogVisible: {
      get() {
        return this.visible
      },
      set(val) {
        this.$emit('update:visible', val)
      }
    },
    //上传接口地址
    uploadAction() {
      return apiUrl + '/UploadImage'
    },
    //上传携带的鉴权头
    uploadHeaders() {
      return {Authorization: localStorage.getItem('token')}
    },
    //图片预览地址：相对路径拼上后端服务地址
    albumUrl() {
      if (!this.form.album) {
        return ''
      }
      if (/^https?:\/\//.test(this.form.album)) {
        return this.form.album
      }
      return apiUrl.replace('/api/boss', '') + this.form.album
    }
  },
  watch: {
    editData: {
      handler(data) {
        if (data && Object.keys(data).length > 0) {
          this.form = {...data}
        }
      }
    }
  },
  methods: {
    //上传前校验
    beforeUpload(file) {
      const isImage = /^image\/(png|jpe?g|gif)$/.test(file.type)
      if (!isImage) {
        this.$message.error('仅支持上传 png/jpg/gif/jpeg 图片')
      }
      const isLt2M = file.size / 1024 / 1024 < 2
      if (!isLt2M) {
        this.$message.error('图片大小不能超过 2MB')
      }
      return isImage && isLt2M
    },
    //上传成功
    onUploadSuccess(res) {
      if (res.code === 0) {
        this.form.album = res.url
        this.$message.success('图片上传成功')
      } else {
        this.$message.error(res.message || '图片上传失败')
      }
    },
    //上传失败
    onUploadError() {
      this.$message.error('图片上传失败，请检查网络连接')
    },
    //添加或修改
    async save() {
      try {
        const addOrSave = this.form.id ? SaveProduct : AddProduct;
        let message = await addOrSave({...this.form})
        this.$message.success(message)
        this.handleClose()
        this.$emit('done')
      } catch (e) {
        this.$message.error(e.message);
      }
    },
    //关闭编辑对话框
    handleClose() {
      this.dialogVisible = false
      this.resetForm()
    },
    //重置表单
    resetForm() {
      Object.keys(this.form).forEach(key => {
        this.form[key] = ''
      })
      if (this.$refs.form) {
        this.$refs.form.clearValidate()
      }
    }
  }
}
</script>

<style scoped>
/* 压缩图片行的占高，减少与下方简介的空白 */
.album-item {
  margin-bottom: 10px;
}

.album-wrap {
  display: flex;
  align-items: center;
}

.album-side {
  margin-left: 12px;
}

.album-uploader >>> .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  overflow: hidden;
  transition: border-color .3s;
}

.album-uploader >>> .el-upload:hover {
  border-color: #409EFF;
}

.album-preview {
  width: 120px;
  height: 120px;
  display: block;
  object-fit: cover;
}

.album-placeholder {
  width: 120px;
  height: 120px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: #8c939d;
}

.album-placeholder i {
  font-size: 28px;
}

.album-placeholder span {
  margin-top: 8px;
  font-size: 13px;
}

.album-tip {
  margin: 0 0 10px;
  font-size: 12px;
  color: #909399;
}
</style>

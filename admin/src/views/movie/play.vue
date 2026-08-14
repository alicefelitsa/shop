<template>
  <div>
    <el-dialog :title="form.title || '播放视频'" :visible.sync="dialogVisible" width="57%"
               class="responsive-dialog" :close-on-click-modal="false" :before-close="handleClose" top="3vh">
      <!-- 视频播放器 -->
      <div ref="dplayerContainer" class="video-wrapper"></div>
      <!-- 集数列表 -->
      <div v-if="episode.length > 0" class="episode-list">
        <div class="episode-buttons">
          <el-button
              v-for="(item, index) in episode"
              :key="index"
              size="mini"
              :type="currentIndex === index ? 'primary' : ''"
              @click="selectEpisode(index)"
          >
            {{ `第${item.page}集` }}
          </el-button>
        </div>
      </div>
      <!--      <span slot="footer" class="dialog-footer">-->
      <!--        <el-button @click="handleClose" size="small">关 闭</el-button>-->
      <!--      </span>-->
    </el-dialog>
  </div>
</template>

<script>
import DPlayer from 'dplayer'
import {GetPlayVideo} from "@/api/movie";

// 播放视频
export default {
  props: {
    //弹窗是否打开
    visible: {
      type: Boolean,
      default: false
    },
    //修改回显的数据
    playData: {
      type: Object,
      default: () => ({})
    }
  },
  data() {
    return {
      form: {
        id: '',
        title: '',
        category: '',
        region: '',
        director: '',
        year: '',
        format: '',
        play_url: '',
        episode: '',
        plot: '',
        score: '',
        cover: ''
      },
      player: null,
      episode: [],
      videoUrl: '',
      currentIndex: 0
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
    }
  },
  watch: {
    playData: {
      handler(data) {
        if (data && Object.keys(data).length > 0) {
          this.form = {...data}
          this.$nextTick(async () => {
            await this.getPlayVideo()
            this.initPlayer()
          })
        }
      }
    }
  },
  beforeDestroy() {
    this.destroyPlayer()
  },
  methods: {
    //获取播放影片
    async getPlayVideo() {
      try {
        let data = await GetPlayVideo({...this.form})
        console.log(data)
        this.episode = data
        this.videoUrl = encodeURI(this.episode[0].playUri)
      } catch (e) {
        this.$message.error(e.message);
      }
    },
    //选择集数
    selectEpisode(index) {
      if (this.currentIndex === index) return
      this.currentIndex = index
      this.videoUrl = encodeURI(this.episode[index].playUri)
      this.$nextTick(() => {
        this.initPlayer()
      })
    },
    //初始化播放器
    initPlayer() {
      this.destroyPlayer()
      const options = {
        container: this.$refs.dplayerContainer,
        autoplay: false,
        theme: '#b7daff',
        loop: false,
        lang: 'zh-cn',
        screenshot: false,
        hotkey: true,
        preload: 'auto',
        volume: 0.5,
        mutex: true,
        video: {
          url: this.videoUrl,
          pic: this.form.cover || '',
          type: 'auto'
        }
      }
      this.player = new DPlayer(options)
      // this.$nextTick(() => {
      //   this.player.volume(0.5, true)
      // })
      // this.player.on('error', () => {
      //   if (this.dialogVisible) {
      //     this.$message.error('视频加载失败')
      //   }
      // })
    },
    //销毁播放器实例
    destroyPlayer() {
      if (this.player) {
        this.player.destroy()
        this.player = null
      }
    },
    //关闭编辑对话框
    handleClose() {
      this.currentIndex = 0
      this.destroyPlayer()
      this.dialogVisible = false
      this.resetForm()
      this.$emit('done')
    },
    //重置表单
    resetForm() {
      Object.keys(this.form).forEach(key => {
        this.form[key] = ''
      })
    }
  }
}
</script>

<style scoped>
.video-wrapper {
  width: 100%;
  background: #000;
  margin-top: -10px;
}

.video-wrapper >>> .dplayer-video-wrap {
  width: 100%;
  height: auto;
  aspect-ratio: 16 / 9;
}

.video-wrapper >>> .dplayer-video-current {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.episode-list {
  margin-top: 15px;
  margin-bottom: 12px;
  border-radius: 4px;
}

.episode-buttons {
  display: flex;
  flex-wrap: wrap;
  gap: 9px;
}

.episode-buttons .el-button {
  margin: 0;
  padding: 7px 10px !important;
}
</style>
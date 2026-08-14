<template>
  <div class="chat-page">
    <!-- 顶部搜索栏 -->
    <!--    <el-header class="page-header">
          <el-input
              v-model="searchQuery"
              placeholder="请输入搜索的访客"
              clearable
              style="width: 200px; margin-right: 10px;"
          >
            <el-button slot="append" :icon="Search" type="primary" size="small"/>
          </el-input>
          <span class="header-title">访客：{{ currentVisitor.id }}【{{ currentVisitor.location }}】</span>
        </el-header>-->

    <el-container class="main-container">
      <!-- 左侧访客列表 -->
      <el-aside class="visitor-sidebar" v-loading="visitorLoading">
        <!-- 标签栏 -->
        <!--        <el-tabs v-model="activeTab" type="card" size="small">
                  <el-tab-pane label="全部" name="all"/>
                  <el-tab-pane label="未读" name="unread">
                    <el-badge :value="0" type="danger" slot="label"/>
                  </el-tab-pane>
                  <el-tab-pane label="拉黑" name="blocked"/>
                </el-tabs>-->

        <!-- 访客列表 -->
        <div
            v-for="visitor in visitorList"
            :key="visitor.uuid"
            class="visitor-item"
            :class="{ 'active': currentVisitor?.uuid === visitor.uuid }"
            @click="selectVisitor(visitor)"
        >
          <el-avatar :size="45" :src="file_url+visitor.cover"/>
          <div class="visitor-info">
            <div class="visitor-header">
              <div class="visitor-id">{{ visitor.uuid }}</div>
              <div class="visitor-time">{{ visitor.online_time }}</div>
            </div>
            <div class="visitor-last-msg">
              <i :class="visitor.status===0 ? 'online-dot' : 'offline-dot'"/>
              <el-badge :hidden="visitor.is_read_num === 0" :value="visitor.is_read_num" class="item">
                <span class="last-msg-text">{{ visitor.lastMsg ?? '无新信息' }}</span>
              </el-badge>
            </div>
            <div class="visitor-meta">
              <span class="positionSpan">{{ visitor.position }}</span>
              <span :class="visitor.status===0 ? 'online-text' : 'offline-text'">
                  {{ visitor.status === 0 ? '在线' : '离线' }}
                </span>
            </div>
          </div>
        </div>
      </el-aside>

      <el-container>
        <template v-if="currentVisitor">
          <!-- 中间聊天区域 -->
          <el-main class="chat-main">
            <div class="chat-header">
              <el-avatar :size="36" :src="currentVisitor.photo"/>
              <span class="chat-header-title">{{ currentVisitor.uuid }}</span>
              <!--          <el-avatar :size="36" :src="adminAvatar" class="admin-avatar"/>-->
            </div>

            <!-- 聊天消息区域 -->
            <div class="chat-messages" ref="messageContainer" v-loading="loading">
              <div
                  v-for="(msg, index) in messageList"
                  :key="index"
                  class="message-item"
                  :class="msg.role===0 ? 'self-message' : 'other-message'"
              >
                <el-avatar
                    v-if="msg.role===1"
                    :size="41"
                    :src="currentVisitor.photo"
                    class="message-avatar"
                />
                <div class="message-content">
                  <div class="message-text" :class="msg.role===0 ? 'self-text' : 'other-text'">
                    <template v-if="msg.types==='text'">
                      <div class="bubble" v-html="msg.content"></div>
                    </template>
                    <template v-if="msg.types==='image'">
                      <el-image
                          style="width: 200px;"
                          :src="file_url+msg.content"
                          :preview-src-list="imageList">
                      </el-image>
                    </template>
                    <template v-if="msg.types==='video'">
                      <video :key="msg.content" controls style="width: 100%;height: 300px;">
                        <source :src="file_url+msg.content" type="video/mp4">
                      </video>
                    </template>
                  </div>
                  <div class="message-time">{{ msg.ctime }}</div>
                  <!--                  <div v-if="msg.role===0" class="message-status">已读</div>-->
                </div>
                <el-avatar
                    v-if="msg.role===0"
                    :size="41"
                    :src="userAvatar"
                    class="message-avatar"
                />
              </div>
            </div>

            <!-- 输入工具栏 -->
            <div class="chat-toolbar">
              <el-tooltip content="发送表情" placement="top">
                <el-popover
                    placement="top-start"
                    width="400"
                    trigger="click">
                  <template>
                    <div class="popper-emoji clearfix">
                      <span v-for="(item, index) in emoji" :key="index" @click="inputEmoji(item)">{{ item }}</span>
                    </div>
                  </template>
                  <i slot="reference" class="iconfont icon-smile"></i>
                </el-popover>
              </el-tooltip>
              <el-tooltip content="发送图片" placement="top">
                <el-upload
                    ref="uploadRef"
                    name="file"
                    :action="file_url+'/uploading/api/uploadImages'"
                    multiple
                    :limit="2"
                    accept="image/*"
                    :show-file-list="false"
                    :on-success="handleAvatarSuccess"
                    :on-exceed="handleExceed"
                    :before-upload="beforeAvatarUpload"
                >
                  <i class="iconfont icon-tupian"></i>
                </el-upload>
              </el-tooltip>
              <!--              <el-tooltip content="发送位置" placement="top">
                              <i class="iconfont icon-location_light"></i>
                            </el-tooltip>-->
              <!--          <el-tooltip content="快捷回复" placement="top">
                          <i class="iconfont icon-commentso"></i>
                        </el-tooltip>-->
              <el-tooltip content="一键撤回" placement="top">
                <i class="iconfont icon-qinglihuancun"></i>
              </el-tooltip>
            </div>

            <!-- 输入框 & 发送按钮 -->
            <div class="chat-input-wrapper">
              <el-input
                  v-model="inputContent"
                  type="textarea"
                  :rows="8"
                  placeholder="请填写发送内容...Enter发送，Shift+Enter换行"
                  resize="none"
                  class="chat-input"
                  @keydown.enter.native="handleSend"
              />
              <div class="send-btn-container">
                <!--            [超链接文本【头部必须包含http://或https://】]-->
                <el-button type="primary" size="medium" @click="handleSend">发送[Enter]</el-button>
              </div>
            </div>
          </el-main>

          <!-- 右侧访客资料 -->
          <el-aside class="visitor-info-sidebar">
            <el-tabs v-model="infoTab">
              <el-tab-pane label="访客资料" name="info"></el-tab-pane>
            </el-tabs>
            <div v-if="infoTab === 'info'" class="info-panel">
              <!-- 头像区域 -->
              <div class="info-item avatar-item">
                <el-avatar :size="60" :src="currentVisitor.photo"/>
              </div>
              <!-- 信息项列表 -->
              <div class="info-item">
                <span class="info-label">访客ID：</span>
                <span class="info-value">{{ currentVisitor.uuid }}</span>
              </div>
              <div class="info-item">
                <span class="info-label">访客昵称：</span>
                <span class="info-value">{{ currentVisitor.nickname }}</span>
              </div>
              <div class="info-item">
                <span class="info-label">扫码次数：</span>
                <span class="info-value">{{ currentVisitor.scan_num }}</span>
              </div>
              <div class="info-item">
                <span class="info-label">注册时间：</span>
                <span class="info-value">{{ currentVisitor.ctime }}</span>
              </div>
              <div class="info-item">
                <span class="info-label">登录时间：</span>
                <span class="info-value">{{ currentVisitor.online_time }}</span>
              </div>
              <div class="info-item">
                <span class="info-label">离线时间：</span>
                <span class="info-value">{{ currentVisitor.offline_time }}</span>
              </div>
              <div class="info-item">
                <span class="info-label">IP地址：</span>
                <span class="info-value">{{ currentVisitor.ip }}</span>
              </div>
              <div class="info-item">
                <span class="info-label">位置：</span>
                <span class="info-value">{{ currentVisitor.position }}</span>
              </div>
              <div class="info-item">
                <span class="info-label">设备：</span>
                <span class="info-value multi-line">{{ currentVisitor.device }}</span>
              </div>
            </div>
          </el-aside>
        </template>
        <template v-else>
          <el-empty description="请选择访客进行聊天" style="width: 100%"></el-empty>
        </template>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import {getChat, getClientInfo, getClientList, saveChat, updateIsReadNum} from "@/api/chat";
import dayjs from "dayjs";

export default {
  name: 'ChatPage',
  data() {
    return {
      infoTab: 'info',
      emoji: [
        "😀",
        "😁",
        "😂",
        "😃",
        "😄",
        "😅",
        "😆",
        "😉",
        "😊",
        "😋",
        "😎",
        "😍",
        "😘",
        "😗",
        "😙",
        "😚",
        "😇",
        "😐",
        "😑",
        "😶",
        "😏",
        "😣",
        "😥",
        "😮",
        "😯",
        "😪",
        "😫",
        "😴",
        "😌",
        "😛",
        "😜",
        "😝",
        "😒",
        "😓",
        "😔",
        "😕",
        "😲",
        "😷",
        "😖",
        "😞",
        "😟",
        "😤",
        "😢",
        "😭",
        "😦",
        "😧",
        "😨",
        "😬",
        "😰",
        "😱",
        "😳",
        "😵",
        "😡",
        "😠",
        "😈",
        "👿",
        "💩",
        "👻",
        "🙌",
        "🖕",
        "👍",
        "👫",
        "👬",
        "👭",
        "🌚",
        "🙈",
        "💊",
        "🙏",
        "🍦",
        "🍉",
        "🐁",
        "🐂",
        "🐅",
        "🐇",
        "🐉",
        "🐍",
        "🐎",
        "🐐",
        "🐒",
        "🐓",
        "🐕",
        "🐖",
      ],
      searchQuery: '',
      activeTab: 'all', // 访客列表标签（全部/未读/拉黑）
      inputContent: '',
      userAvatar: localStorage.getItem("avatar"), // 客服头像
      visitorList: [],
      currentVisitor: null,
      messageList: [],
      file_url: '',
      upload_url: 'http://38.76.211.75',
      ws_address: '',
      visitorLoading: false,
      loading: false,
      imageList: [],
      ucode: localStorage.getItem("ucode"),
      uploadSuccessNum: 0,
      ws: null,
      wsHeartbeatTimer: null,
      wsReconnectTimer: null,
      wsRetryCount: 0,
      wsMaxRetry: 10,
      wsHeartbeatInterval: 25000, // 25s
      wsIsManualClose: false
    }
  },
  mounted() {
    this.getVisitorList();
    document.addEventListener('visibilitychange', this.onVisibilityChange);
    window.addEventListener('online', this.onNetworkOnline);
  },
  beforeDestroy() {
    this.wsIsManualClose = true;
    this.stopHeartbeat();
    clearTimeout(this.wsReconnectTimer);
    this.ws?.close(1000, '页面卸载');
  },
  methods: {
    //连接WebSocket
    initWebSocket() {
      if (!this.ucode) return;

      // ✅ 防止重复连接
      if (this.ws && [WebSocket.OPEN, WebSocket.CONNECTING].includes(this.ws.readyState)) {
        return;
      }

      this.wsIsManualClose = false;
      const wsUrl = `${this.ws_address}/user?ucode=${encodeURIComponent(this.ucode)}`;
      this.ws = new WebSocket(wsUrl);

      /** ✅ 连接成功 */
      this.ws.onopen = () => {
        console.log('✅WebSocket连接成功');
        this.wsRetryCount = 0;
        this.startHeartbeat();
      };

      /** ✅ 接收消息 */
      this.ws.onmessage = (event) => {
        try {
          const data = JSON.parse(event.data);
          // 放入微任务队列，不阻塞主线程
          queueMicrotask(() => {
            try {
              this.handleWsMessage(data);
            } catch (error) {
              console.error('处理消息时出错:', error);
            }
          });
        } catch (error) {
          console.error('解析消息失败:', error, '原始数据:', event.data);
        }
      };

      /** ✅ 错误 */
      this.ws.onerror = (err) => {
        console.error('❌WebSocket 错误:', err);
      };

      /** ✅ 关闭 → 自动重连 */
      this.ws.onclose = () => {
        console.warn('⚠️ WebSocket关闭');
        this.stopHeartbeat();
        if (!this.wsIsManualClose) {
          this.reconnect();
        }
      };
    },
    //重新连接WS
    reconnect() {
      if (this.wsRetryCount >= this.wsMaxRetry) {
        console.error('❌ WebSocket 重连失败次数过多');
        return;
      }
      const delay = Math.min(2000 * 2 ** this.wsRetryCount, 30000);
      this.wsRetryCount++;
      console.log(`🔄 WebSocket 重连中（${this.wsRetryCount}）...`);
      this.wsReconnectTimer = setTimeout(() => {
        this.initWebSocket();
      }, delay);
    },
    //开始心跳
    startHeartbeat() {
      this.stopHeartbeat();
      this.wsHeartbeatTimer = setInterval(() => {
        if (this.ws?.readyState === WebSocket.OPEN) {
          this.ws.send(JSON.stringify({type: 'ping'}));
          //this.ws.ping('heartbeat');
        }
      }, this.wsHeartbeatInterval);
    },
    //停止心跳
    stopHeartbeat() {
      clearInterval(this.wsHeartbeatTimer);
      this.wsHeartbeatTimer = null;
    },
    onVisibilityChange() {
      if (document.visibilityState === 'visible') {
        this.checkAndReconnect();
      }
    },
    onNetworkOnline() {
      this.checkAndReconnect();
    },
    checkAndReconnect() {
      if (!this.ws || this.ws.readyState !== WebSocket.OPEN) {
        this.initWebSocket();
      }
    },
    //处理接收到的消息
    async handleWsMessage(data) {
      console.log(data);
      let uuid = data.from
      let content = data.content
      if (data.types === 'online') {
        if (this.visitorList.find(item => item.uuid === uuid)) {
          //修改访客上线
          this.visitorList = this.visitorList.map((visitor, index) => {
            if (visitor.uuid === uuid) {
              return {
                ...visitor,
                status: 0,
                online_time: data.ctime
              };
            }
            return visitor;
          }).sort((a, b) => {
            if (a.uuid === uuid) return -1;
            if (b.uuid === uuid) return 1;
            return 0;
          });
        } else {
          //插入新登录的访客资料
          getClientInfo(localStorage.getItem("ucode"), uuid).then((data) => {
            this.visitorList.unshift(data);
          }).catch((e) => {
            this.$message.error(e.message);
          });
        }
      } else if (data.types === 'offline') {
        //修改访客离线
        if (this.visitorList.find(item => item.uuid === uuid)) {
          //修改访客上线
          this.visitorList = this.visitorList.map((visitor, index) => {
            if (visitor.uuid === uuid) {
              return {
                ...visitor,
                status: 1
              };
            }
            return visitor;
          }).sort((a, b) => a.status - b.status);
        }
      } else {
        if (data.types === 'text') {
          if (this.currentVisitor?.uuid === uuid) {
            this.messageList.push(data)
            this.$nextTick(() => {
              this.$refs.messageContainer.scrollTop = this.$refs.messageContainer.scrollHeight
            })
          }
        }
        if (data.types === 'image') {
          if (this.currentVisitor?.uuid === uuid) {
            this.messageList.push(data)
            this.imageList.push(this.file_url + content);
            setTimeout(() => {
              this.$nextTick(() => {
                this.$refs.messageContainer.scrollTop = this.$refs.messageContainer.scrollHeight
              })
            }, 1000)
          }
        }
        //更新访客发送的最后一条信息
        if (this.visitorList.find(item => item.uuid === uuid)) {
          this.visitorList = this.visitorList.map((visitor, index) => {
            if (visitor.uuid === uuid) {
              return {
                ...visitor,
                lastMsg: content,
                is_read_num: this.currentVisitor?.uuid !== uuid ? visitor.is_read_num + 1 : 0
              };
            }
            return visitor;
          })
        }
        //选中的访客有消息接收更新数据库未读为0
        if (this.currentVisitor?.uuid === uuid) {
          setTimeout(async () => {
            let msg = await updateIsReadNum(uuid)
            //this.$message.success(msg)
          }, 500)
        }
      }
    },
    // 发送消息到WS
    sendWs(msg) {
      if (this.ws?.readyState === WebSocket.OPEN) {
        this.ws.send(JSON.stringify(msg));
      } else {
        console.warn('⚠️ WebSocket未连接，稍后重试');
        this.checkAndReconnect();
      }
    },
    //获取访客列表
    getVisitorList() {
      this.visitorLoading = true;
      getClientList(localStorage.getItem("ucode")).then((data) => {
        //console.log(data);
        this.visitorList = data.client;
        this.file_url = data.config.file_url;
        this.ws_address = data.config.ws_address;
        this.initWebSocket();
        setTimeout(() => {
          this.visitorLoading = false;
          this.messageList = data;
        }, 200)
      }).catch((e) => {
        this.$message.error(e.message);
      });
    },
    // 点击选择访客
    selectVisitor(visitor) {
      this.loading = true;
      this.currentVisitor = visitor
      this.currentVisitor.is_read_num = 0
      this.currentVisitor.photo = this.file_url + visitor.cover
      this.messageList = [];
      this.imageList = []
      getChat(localStorage.getItem("ucode"), this.currentVisitor.uuid).then((data) => {
        setTimeout(() => {
          this.loading = false;
          this.messageList = data;
          data.forEach((item, index) => {
            if (item.types === 'image') {
              this.imageList.push(this.file_url + item.content);
            }
          });
        }, 200)
        //滚动到底部
        setTimeout(() => {
          this.$nextTick(() => {
            this.$refs.messageContainer.scrollTop = this.$refs.messageContainer.scrollHeight
          })
        }, 1000)
      }).catch((e) => {
        this.$message.error(e.message);
      });
    },
    //插入表情
    inputEmoji(emoji) {
      this.inputContent += emoji;
    },
    //图片上传成功
    handleAvatarSuccess(response, file, fileList) {
      if (response.code === 0) {
        this.uploadSuccessNum++
        let image = response.filePath
        this.imageList.push(this.file_url + image);
        let newMsg = {
          role: 0,
          content: image,
          ctime: dayjs().format('YYYY-MM-DD HH:mm:ss'),
          types: 'image',
          from: localStorage.getItem("ucode"),
          to: this.currentVisitor.uuid,
        }
        this.messageList.push(newMsg)
        this.sendWs(newMsg);
        this.recordChat(newMsg);
        if (this.uploadSuccessNum === fileList.length) {
          //清空此次上传完成的图片数量
          this.$nextTick(() => {
            this.$refs.uploadRef.clearFiles()
          })
          this.uploadSuccessNum = 0
        }
        //滚动到底部
        setTimeout(() => {
          this.$nextTick(() => {
            this.$refs.messageContainer.scrollTop = this.$refs.messageContainer.scrollHeight
          })
        }, 1000)
      } else {
        this.$message.error(response.message)
      }
    },
    handleExceed(files, fileList) {
      this.$message.warning(`当前限制选择 2 个文件，本次选择了 ${files.length} 个文件。`);
    },
    //上传图片前钩子
    beforeAvatarUpload(file) {
      const size = file.size / 1024 / 1024 < 3;
      if (!size) {
        this.$message.error('上传图片大小不能超过 3MB');
        return false
      }
    },
    //保存聊天信息
    recordChat(newMsg) {
      saveChat(newMsg).then((msg) => {
        console.log(msg);
        //this.$message.success(msg);
      }).catch((e) => {
        this.$message.error(e.message);
      });
    },
    // 发送消息
    handleSend(event) {
      if (event.shiftKey) {
        return; // 返回，不阻止默认行为
      }
      event.preventDefault();
      if (!this.inputContent.trim()) return
      let newMsg = {
        role: 0,
        content: this.inputContent.trim(),
        ctime: dayjs().format('YYYY-MM-DD HH:mm:ss'),
        types: 'text',
        from: localStorage.getItem("ucode"),
        to: this.currentVisitor.uuid,
      }
      this.messageList.push(newMsg)
      this.inputContent = ''
      this.sendWs(newMsg);
      this.recordChat(newMsg);
      // 滚动到底部
      this.$nextTick(() => {
        this.$refs.messageContainer.scrollTop = this.$refs.messageContainer.scrollHeight
      })
    }
  }
}
</script>

<style>
/* 全局容器 */
.chat-page {
  background-color: #f5f7fa;
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", Arial, sans-serif;
  height: 100vh; /* 确保整个页面占满视口高度 */
  overflow: hidden;
  width: 100vw;
}

/* 顶部搜索栏 */
.page-header {
  display: flex;
  align-items: center;
  background-color: #409eff;
  color: #fff;
  padding: 0 20px;
  height: 60px;
}

.header-title {
  margin-left: 20px;
  font-size: 16px;
  font-weight: bold;
}

/* 主容器（左右布局） */
.main-container {
  display: flex;
  height: 100%; /* 继承父级 100vh */
  width: 100%;
  /*height: 97vh;*/
}

.main-container > .el-container {
  display: flex;
  flex-direction: row;
  overflow: hidden;
  flex: 1;
}

/* 左侧访客列表 */
.visitor-sidebar {
  background-color: #fff;
  border-right: 1px solid #e6e6e6;
  /*border-bottom: 1px solid #e6e6e6;*/
  display: flex;
  flex-direction: column;
  width: 290px !important;
}

.visitor-list-scroll {
  flex: 1;
  overflow-y: auto;
}

.visitor-item {
  display: flex;
  align-items: center;
  padding: 10px;
  cursor: pointer;
  border-bottom: 1px solid #f0f0f0;
  transition: background-color 0.2s;
}

.visitor-item:hover {
  background-color: #f5f7fa;
}

.visitor-item.active {
  background-color: #ecf5ff;
}

.visitor-info {
  margin-left: 10px;
  flex: 1;
}

.visitor-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 4px; /* 添加一点底部间距 */
}

.visitor-time {
  font-size: 12px;
  color: #909399;
}

.visitor-id {
  font-weight: bold;
  font-size: 14px;
  color: #8c7398;
}

.visitor-last-msg {
  font-size: 12px;
  color: #606266;
  display: flex;
  align-items: center;
  height: 20px;
}

.last-msg-text {
  /* 核心属性：强制一行显示 */
  white-space: nowrap;
  /* 核心属性：隐藏超出部分 */
  overflow: hidden;
  /* 核心属性：超出显示 ... */
  text-overflow: ellipsis;
  /* 可选：设置最大宽度，根据实际布局调整 */
  /* 例如限制最大宽度为 120px，或者用 flex 布局自动收缩 */
  width: 160px;
  display: inline-block; /* 为了让 max-width 生效 */
  margin-right: 11px;
}

.positionSpan {
  /* 核心属性：强制一行显示 */
  white-space: nowrap;
  /* 核心属性：隐藏超出部分 */
  overflow: hidden;
  /* 核心属性：超出显示 ... */
  text-overflow: ellipsis;
  /* 可选：设置最大宽度，根据实际布局调整 */
  /* 例如限制最大宽度为 120px，或者用 flex 布局自动收缩 */
  width: 170px;
}

.online-dot {
  display: inline-block;
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background-color: #67c23a;
  margin-right: 4px;
}

.offline-dot {
  display: inline-block;
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background-color: #909399;
  margin-right: 4px;
}

.visitor-meta {
  display: flex;
  justify-content: space-between;
  font-size: 12px;
  color: #909399;
  margin-top: 4px;
}

.online-text {
  color: #67c23a;
}

.offline-text {
  color: #909399;
}

/* 中间聊天区域 */
.chat-main {
  flex: 1;
  display: flex;
  flex-direction: column;
  background-color: #fff;
  padding: 0;
  overflow: hidden;
  min-width: 500px;
}

.chat-header {
  display: flex;
  align-items: center;
  padding: 10px 20px;
  border-bottom: 1px solid #e6e6e6;
  background-color: #fafafa;
}

.chat-header-title {
  margin: 0 10px;
  font-weight: bold;
  font-size: 16px;
  color: #303133;
}

.admin-avatar {
  margin-left: auto;
}

.chat-messages {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
  background-color: #f5f7fa;
  font-size: 16px;
}

.message-item {
  display: flex;
  margin-bottom: 20px;
  align-items: flex-start;
}

.self-message {
  display: flex;
  flex-direction: row; /* 修改：原来是 row-reverse，改为正向排列 */
  justify-content: flex-end; /* 修改：整体靠右对齐 */
  margin-bottom: 20px;
  align-items: flex-start; /* 新增：确保头像和文字顶部对齐 */
}

.message-avatar {
  margin: 0 10px;
}

.message-content {
  max-width: 70%;
  display: flex;
  flex-direction: column;
}

.message-text {
  line-height: 26px;
  white-space: pre-wrap;
  word-wrap: break-word; /* 长单词或URL换行 */
  word-break: break-word; /* 中文换行 */
}

.self-text {
  background-color: #409eff;
  color: #fff;
  border-radius: 8px 8px 0 8px;
  padding: 8px;
  position: relative;
}

.other-text {
  background-color: #fff;
  color: #303133;
  border: 1px solid #e6e6e6;
  border-radius: 8px 8px 8px 0;
  padding: 8px;
}

.message-time {
  font-size: 12px;
  color: #909399;
  margin-top: 4px;
  text-align: left;
}

.self-message .message-time {
  text-align: right;
}

.message-status {
  font-size: 12px;
  color: #67c23a;
  margin-top: 2px;
  text-align: right;
}

.self-message .message-status {
  text-align: left;
}

/* 输入工具栏 */
.chat-toolbar {
  display: flex;
  align-items: center;
  padding: 0 20px;
  border-top: 1px solid #e6e6e6;
  background-color: #fafafa;
}

.chat-toolbar .el-button {
  margin-right: 10px;
  color: #409eff;
}

/* 输入框 & 发送提示 */
.chat-input {
  margin: 0 0 0 0;
  font-size: 16px !important;
}

.send-tips {
  font-size: 12px;
  color: #67c23a;
  margin: 10px 20px 10px;
  line-height: 1.4;
}

/* 右侧访客资料 */
.visitor-info-sidebar {
  min-width: 320px !important;
  flex-shrink: 0; /* 防止被压缩 */
  border-left: 1px solid #e6e6e6;
  padding: 17px 20px 0 20px;
  overflow-y: auto;
  overflow-x: hidden;
  background-color: #fff;
}

.popper-emoji {
  span {
    float: left;
    width: 33px;
    height: 33px;
    line-height: 33px;
    text-align: center;
    font-size: 22px;
    cursor: pointer;
    border-radius: 3px;
    overflow: hidden;
  }
}

.iconfont {
  display: flex;
  width: 40px;
  height: 40px;
  justify-content: center;
  align-items: center;
  font-size: 24px;
  cursor: pointer;
}

/* 移除焦点状态下的边框和阴影 */
.chat-input ::v-deep .el-textarea__inner:focus {
  border: 1px solid #DCDFE6;
  outline: none;
  box-shadow: none;
}

/* 如果上面不行，可以使用更深的穿透 */
.chat-input :deep(.el-textarea__inner:focus) {
  border: 1px solid #DCDFE6;
  outline: none;
  box-shadow: none;
}

.chat-input-wrapper {
  position: relative;
  width: 100%;
}

.chat-input {
  font-size: 14px;
  width: 100%;
  padding-right: 0; /* 为按钮留出空间 */
}

.el-textarea__inner {
  border-top: 1px solid #e6e6e6 !important;
  border-bottom: 0 !important;
  border-left: 0 !important;
  border-right: 0 !important;
  border-radius: 0 !important;
  box-shadow: none !important;
}


.send-btn-container {
  position: absolute;
  right: 10px;
  bottom: 10px;
  z-index: 10; /* 确保按钮在输入框上方 */
  font-size: 14px;
  color: #0da018;
}

.info-panel {
  display: flex;
  flex-direction: column;
  align-items: center; /* 内容居中 */
  padding: 16px;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  max-width: 400px; /* 限制最大宽度，避免过宽 */
  margin: 0 auto; /* 页面居中（可选） */
}

/* 头像区域 */
.avatar-item {
  margin-bottom: 16px;
}

/* 信息项容器 */
.info-item {
  width: 100%;
  display: flex;
  justify-content: space-between; /* 标签左对齐，值右对齐（或根据需求调整） */
  align-items: center;
  margin-bottom: 8px;
  font-size: 14px; /* 基础字体大小 */
  color: #333; /* 文字主色 */
}

/* 标签样式 */
.info-label {
  font-weight: bold; /* 标签加粗 */
  color: #666; /* 标签颜色稍浅 */
  margin-right: 8px;
  white-space: nowrap; /* 标签不换行 */
}

/* 值的样式 */
.info-value {
  flex: 1; /* 让值占据剩余空间 */
  text-align: left; /* 值左对齐（或 right 右对齐） */
  color: #333; /* 文字主色 */
  line-height: 1.5; /* 行高，提升可读性 */
  overflow: hidden; /* 隐藏溢出内容 */
  text-overflow: ellipsis; /* 溢出显示省略号（单行） */
  white-space: nowrap; /* 单行模式（默认） */
}

/* 多行模式（针对长文本，如位置、系统） */
.info-value.multi-line {
  white-space: pre-wrap; /* 保留换行和空格，自动换行 */
  word-break: break-all; /* 长单词/URL 强制换行 */
  text-overflow: clip; /* 取消省略号，允许换行 */
  max-height: none; /* 取消高度限制 */
}

/* 在线状态颜色 */
.online-text {
  color: #00b42a; /* 在线绿色 */
}

.offline-text {
  color: #ff4d4f; /* 离线红色 */
}

.bubble p {
  margin: 0;
  animation: fadeIn 0.3s ease;
  line-height: 26px;
  word-wrap: break-word; /* 长单词换行 */
  overflow-wrap: break-word; /* 更现代的属性 */
  white-space: pre-wrap; /* 保留空格和换行，但自动换行 */
}

.el-badge__content.is-fixed {
  top: 8px
}

/* 方案3：渐隐渐现的极细滚动条 */
.chat-messages::-webkit-scrollbar {
  width: 6px !important;
  height: 2px !important;
}

.chat-messages::-webkit-scrollbar-track {
  background: linear-gradient(to bottom,
  transparent 0%,
  rgba(0, 0, 0, 0.02) 20%,
  rgba(0, 0, 0, 0.02) 80%,
  transparent 100%) !important;
  border-radius: 0 !important;
}

.chat-messages::-webkit-scrollbar-thumb {
  background: linear-gradient(to bottom,
  transparent 0%,
  rgba(0, 0, 0, 0.15) 30%,
  rgba(0, 0, 0, 0.2) 50%,
  rgba(0, 0, 0, 0.15) 70%,
  transparent 100%) !important;
  border-radius: 0 !important;
  min-height: 40px !important; /* 确保滑块有最小高度 */
}

</style>
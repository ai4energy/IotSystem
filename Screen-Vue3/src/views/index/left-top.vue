<script setup lang="ts">
import { reactive, ref, onMounted, onUnmounted } from "vue";
import { countDeviceNum } from "@/api";
import CountUp from "@/components/count-up";
import { ElMessage } from "element-plus"

const duration = ref(2);
const state = reactive({
  alarmNum: 0,
  offlineNum: 0,
  onlineNum: 0,
  totalNum: 0,
});


const getData = async () => {
  try {
    const response = await fetch('/api/sensors'); // 替换为你的实际 API 地址
    if (!response.ok) {
      throw new Error('网络响应不正常');
    }
    const data = await response.json();
    if (data) {
      console.log("获取数据成功:", data);
      state.alarmNum = data[0].value;
      state.offlineNum = data[1].value;
      state.onlineNum = data[2].value;
      state.totalNum = data[3].value;
    } else {
      ElMessage.error('获取数据失败');
    }
  } catch (error) {
    console.error('请求出错:', error);
    ElMessage.error('请求出错，请稍后再试');
  }
};

let intervalId: NodeJS.Timeout | null = null;

onMounted(() => {
  intervalId = setInterval(getData, 1000); // 每隔 1 秒执行一次 getData
  getData(); // 初始化时立即执行一次
});

onUnmounted(() => {
  if (intervalId) {
    clearInterval(intervalId); // 清除定时器，防止内存泄漏
  }
});
</script>

<template>
  <ul class="user_Overview flex">
    <li class="user_Overview-item" style="color: #00fdfa">
      <div class="user_Overview_nums allnum">
        <CountUp :endVal="state.totalNum" :duration="duration" />
      </div>
      <p>总设备数</p>
    </li>
    <li class="user_Overview-item" style="color: #07f7a8">
      <div class="user_Overview_nums online">
        <CountUp :endVal="state.onlineNum" :duration="duration" />
      </div>
      <p>在线数</p>
    </li>
    <li class="user_Overview-item" style="color: #e3b337">
      <div class="user_Overview_nums offline">
        <CountUp :endVal="state.offlineNum" :duration="duration" />
      </div>
      <p>掉线数</p>
    </li>
    <li class="user_Overview-item" style="color: #f5023d">
      <div class="user_Overview_nums laramnum">
        <CountUp :endVal="state.alarmNum" :duration="duration" />
      </div>
      <p>告警次数</p>
    </li>
  </ul>
</template>

<style scoped lang="scss">
.left-top {
  width: 100%;
  height: 100%;
}

.user_Overview {
  li {
    flex: 1;

    p {
      text-align: center;
      height: 16px;
      font-size: 16px;
    }

    .user_Overview_nums {
      width: 100px;
      height: 100px;
      text-align: center;
      line-height: 100px;
      font-size: 22px;
      margin: 50px auto 30px;
      background-size: cover;
      background-position: center center;
      position: relative;

      &::before {
        content: "";
        position: absolute;
        width: 100%;
        height: 100%;
        top: 0;
        left: 0;
      }

      &.bgdonghua::before {
        animation: rotating 14s linear infinite;
      }
    }

    .allnum {
      &::before {
        background-image: url("@/assets/img/left_top_lan.png");
      }
    }

    .online {
      &::before {
        background-image: url("@/assets/img/left_top_lv.png");
      }
    }

    .offline {
      &::before {
        background-image: url("@/assets/img/left_top_huang.png");
      }
    }

    .laramnum {
      &::before {
        background-image: url("@/assets/img/left_top_hong.png");
      }
    }
  }
}
</style>

<template>
  <div class="login-container">
    <div class="login">
      <el-form style="max-width: 600px" status-icon label-width="auto" class="demo-ruleForm">
        <el-form-item label="用户名" prop="pass">
          <el-input type="password" autocomplete="off" />
        </el-form-item>
        <el-form-item label="密码" prop="checkPass">
          <el-input type="password" autocomplete="off" />
        </el-form-item>
        <el-form-item label="验证码" prop="checkPass">
          <el-input type="password" autocomplete="off" />
          <img src="" alt="" srcset="" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary">登录</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<script setup></script>
<style scoped lang="less">
.login-container {
  width: 100%;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>

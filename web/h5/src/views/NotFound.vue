<template>
  <div class="not-found-page">
    <div class="nf-content container">
      <div class="nf-number">404</div>
      <h1 class="nf-title">Page Not Found</h1>
      <p class="nf-desc">The page you're looking for doesn't exist or has been moved.</p>
      <div class="nf-actions">
        <router-link to="/" class="btn btn-primary btn-lg">Back to Home</router-link>
        <router-link to="/products" class="btn btn-outline btn-lg">Browse Products</router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'NotFound'
}
</script>

<style scoped>
.not-found-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: var(--header-height) 0 0;
  background: var(--bg-light);
}

.nf-content {
  text-align: center;
  padding: 60px 20px;
}

.nf-number {
  font-size: 8rem;
  font-weight: 800;
  color: var(--primary);
  line-height: 1;
  margin-bottom: 16px;
  background: linear-gradient(135deg, var(--primary), var(--accent));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.nf-title {
  font-size: 2rem;
  font-weight: 700;
  color: var(--primary);
  margin-bottom: 12px;
}

.nf-desc {
  font-size: 1.05rem;
  color: var(--text-secondary);
  margin-bottom: 36px;
}

.nf-actions {
  display: flex;
  justify-content: center;
  gap: 16px;
  flex-wrap: wrap;
}
</style>

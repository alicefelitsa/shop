<template>
  <div class="about-page">
    <!-- Page Header -->
    <section class="page-header">
      <div class="page-header-bg"></div>
      <div class="page-header-overlay"></div>
      <div class="container page-header-content">
        <h1 class="page-title">About Us</h1>
        <p class="page-breadcrumb">
          <router-link to="/">Home</router-link>
          <span>/</span>
          <span>About Us</span>
        </p>
      </div>
    </section>

    <!-- Company Story -->
    <section class="section story-section">
      <div class="container">
        <div class="story-grid">
          <div class="story-image">
            <img
              src="@/assets/a1.jpg"
              alt="Our laboratory"
              loading="lazy"
            />
          </div>
          <div class="story-content">
            <span class="section-label">Our Story</span>
            <h2 class="story-title">A Decade of Excellence in Research Supplies</h2>
            <p class="story-text">
              Founded with a vision to support scientific advancement, HKRoids has grown from a small
              laboratory supplier into a globally recognized provider of premium research peptides
              and laboratory equipment. Over the past decade, we have built lasting partnerships
              with research institutions, universities, and pharmaceutical companies worldwide.
            </p>
            <p class="story-text">
              Our state-of-the-art facilities employ rigorous quality control measures, ensuring
              every product meets the highest standards of purity and reliability. Each batch
              undergoes comprehensive HPLC testing, and we provide detailed Certificates of Analysis
              (COA) with every order.
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- Mission & Vision -->
    <section class="section mission-section">
      <div class="container">
        <div class="mission-grid">
          <div class="mission-card">
            <div class="mission-icon">🎯</div>
            <h3 class="mission-title">Our Mission</h3>
            <p class="mission-text">
              To empower researchers and scientists with the highest quality peptides and laboratory
              supplies, enabling breakthrough discoveries that advance human health and scientific
              knowledge. We are committed to accessibility, reliability, and uncompromising quality.
            </p>
          </div>
          <div class="mission-card">
            <div class="mission-icon">🔭</div>
            <h3 class="mission-title">Our Vision</h3>
            <p class="mission-text">
              To become the world's most trusted partner in research peptide supply, setting the
              global standard for quality, innovation, and customer service in the scientific
              supply industry. We envision a future where no research is limited by material quality.
            </p>
          </div>
          <div class="mission-card">
            <div class="mission-icon">💎</div>
            <h3 class="mission-title">Our Values</h3>
            <p class="mission-text">
              Integrity, innovation, and customer-first thinking guide everything we do. We believe
              in transparent business practices, continuous improvement, and building relationships
              that extend far beyond a single transaction.
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- Timeline -->
    <section class="section timeline-section">
      <div class="container">
        <h2 class="section-title">Our Journey</h2>
        <p class="section-subtitle">Key milestones that have shaped our growth and commitment to excellence.</p>
        <div class="timeline">
          <div v-for="(item, idx) in timeline" :key="idx" class="timeline-item" :class="{ 'timeline-right': idx % 2 !== 0 }">
            <div class="timeline-dot"></div>
            <div class="timeline-card">
              <span class="timeline-year">{{ item.year }}</span>
              <h4 class="timeline-title">{{ item.title }}</h4>
              <p class="timeline-desc">{{ item.description }}</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Team / Certifications -->
    <section class="section team-section">
      <div class="container">
        <h2 class="section-title">Our Certifications & Standards</h2>
        <p class="section-subtitle">
          We maintain the highest industry certifications and quality standards.
        </p>
        <div class="cert-grid">
          <div v-for="(cert, idx) in certifications" :key="idx" class="cert-card">
            <div class="cert-icon">{{ cert.icon }}</div>
            <h4 class="cert-title">{{ cert.title }}</h4>
            <p class="cert-desc">{{ cert.description }}</p>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'AboutPage',
  data() {
    return {
      timeline: [
        {
          year: '2014',
          title: 'Company Founded',
          description: 'HKRoids was established with a focus on providing high-quality research peptides to local laboratories and universities.'
        },
        {
          year: '2016',
          title: 'International Expansion',
          description: 'Expanded operations to serve researchers across Europe, North America, and Australia with reliable international shipping.'
        },
        {
          year: '2018',
          title: 'Lab Equipment Division',
          description: 'Launched our laboratory equipment division, offering freeze dryers, synthesis machines, and analytical instruments.'
        },
        {
          year: '2020',
          title: 'Quality Certification',
          description: 'Achieved ISO 9001 quality management certification and established comprehensive HPLC testing protocols for all products.'
        },
        {
          year: '2022',
          title: '5000+ Clients Milestone',
          description: 'Reached over 5,000 satisfied clients worldwide, with a 98% customer satisfaction rate and growing community.'
        },
        {
          year: '2024',
          title: 'Next-Gen Peptide Blends',
          description: 'Introduced proprietary peptide blend formulations and expanded our raw powder catalog to over 60 products.'
        }
      ],
      certifications: [
        {
          icon: '🏅',
          title: 'ISO 9001 Certified',
          description: 'Our quality management system meets international ISO 9001 standards for consistent product excellence.'
        },
        {
          icon: '🔬',
          title: 'HPLC Tested',
          description: 'Every batch undergoes High Performance Liquid Chromatography testing to verify purity above 99%.'
        },
        {
          icon: '📋',
          title: 'COA Provided',
          description: 'Detailed Certificate of Analysis accompanies every order with full testing documentation.'
        },
        {
          icon: '🧬',
          title: 'GMP Compliant',
          description: 'Our manufacturing processes follow Good Manufacturing Practice guidelines for research materials.'
        }
      ]
    }
  }
}
</script>

<style scoped>
/* ===== Page Header ===== */
.page-header {
  position: relative;
  height: 320px;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  padding-top: var(--header-height);
}

.page-header-bg {
  position: absolute;
  inset: 0;
  background-size: cover;
  background-position: center;
  background-color: var(--primary-dark);
  background-image: url('~@/assets/about-bg.svg');
}

.page-header-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(135deg, rgba(15, 36, 64, 0.5) 0%, rgba(26, 54, 93, 0.35) 100%);
}

.page-header-content {
  position: relative;
  z-index: 2;
  text-align: center;
}

.page-title {
  font-size: 2.8rem;
  font-weight: 800;
  color: #fff;
  margin-bottom: 12px;
}

.page-breadcrumb {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  font-size: 0.95rem;
  color: rgba(255, 255, 255, 0.7);
}

.page-breadcrumb a {
  color: var(--accent);
}

.page-breadcrumb a:hover {
  color: var(--accent-light);
}

/* ===== Story Section ===== */
.story-section {
  background: var(--bg-white);
}

.story-grid {
  display: grid;
  grid-template-columns: 1fr 1.2fr;
  gap: 60px;
  align-items: center;
}

.story-image img {
  width: 100%;
  height: 420px;
  object-fit: cover;
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-lg);
}

.section-label {
  display: inline-block;
  font-size: 0.8rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 2px;
  color: var(--accent-dark);
  margin-bottom: 12px;
}

.story-title {
  font-size: 2rem;
  font-weight: 800;
  color: var(--primary);
  line-height: 1.25;
  margin-bottom: 20px;
}

.story-text {
  font-size: 1rem;
  color: var(--text-secondary);
  line-height: 1.8;
  margin-bottom: 16px;
}

/* ===== Mission Section ===== */
.mission-section {
  background: var(--bg-light);
}

.mission-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 28px;
}

.mission-card {
  background: var(--bg-white);
  padding: 40px 32px;
  border-radius: var(--radius-md);
  box-shadow: var(--shadow-sm);
  text-align: center;
  transition: all 0.35s ease;
  border: 1px solid var(--border-light);
}

.mission-card:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow-md);
}

.mission-icon {
  font-size: 2.8rem;
  margin-bottom: 20px;
}

.mission-title {
  font-size: 1.2rem;
  font-weight: 700;
  color: var(--primary);
  margin-bottom: 14px;
}

.mission-text {
  font-size: 0.92rem;
  color: var(--text-secondary);
  line-height: 1.7;
}

/* ===== Timeline ===== */
.timeline-section {
  background: var(--bg-white);
}

.timeline {
  position: relative;
  max-width: 800px;
  margin: 0 auto;
  padding: 20px 0;
}

.timeline::before {
  content: '';
  position: absolute;
  top: 0;
  bottom: 0;
  left: 50%;
  width: 3px;
  background: linear-gradient(to bottom, var(--accent), var(--primary));
  transform: translateX(-50%);
  border-radius: 3px;
}

.timeline-item {
  position: relative;
  padding: 0 0 40px;
  width: 50%;
  padding-right: 48px;
}

.timeline-item.timeline-right {
  margin-left: 50%;
  padding-right: 0;
  padding-left: 48px;
}

.timeline-dot {
  position: absolute;
  top: 8px;
  right: -8px;
  width: 16px;
  height: 16px;
  border-radius: 50%;
  background: var(--accent);
  border: 3px solid var(--bg-white);
  box-shadow: 0 0 0 3px var(--accent);
  z-index: 2;
}

.timeline-item.timeline-right .timeline-dot {
  right: auto;
  left: -8px;
}

.timeline-card {
  background: var(--bg-white);
  padding: 24px;
  border-radius: var(--radius-md);
  border: 1px solid var(--border-color);
  box-shadow: var(--shadow-sm);
  transition: all 0.3s ease;
}

.timeline-card:hover {
  box-shadow: var(--shadow-md);
}

.timeline-year {
  display: inline-block;
  font-size: 0.78rem;
  font-weight: 700;
  background: var(--accent);
  color: var(--primary-dark);
  padding: 3px 12px;
  border-radius: 20px;
  margin-bottom: 10px;
}

.timeline-title {
  font-size: 1.05rem;
  font-weight: 700;
  color: var(--primary);
  margin-bottom: 8px;
}

.timeline-desc {
  font-size: 0.88rem;
  color: var(--text-secondary);
  line-height: 1.6;
}

/* ===== Certifications ===== */
.team-section {
  background: var(--bg-light);
}

.cert-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
}

.cert-card {
  background: var(--bg-white);
  padding: 32px 24px;
  border-radius: var(--radius-md);
  text-align: center;
  border: 1px solid var(--border-color);
  transition: all 0.35s ease;
}

.cert-card:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow-md);
  border-color: transparent;
}

.cert-icon {
  font-size: 2.4rem;
  margin-bottom: 16px;
}

.cert-title {
  font-size: 1rem;
  font-weight: 700;
  color: var(--primary);
  margin-bottom: 10px;
}

.cert-desc {
  font-size: 0.85rem;
  color: var(--text-secondary);
  line-height: 1.6;
}

/* 压缩各区块上下内边距，减少衔接处空白 */
.story-section,
.mission-section,
.timeline-section,
.team-section {
  padding-top: 36px;
  padding-bottom: 36px;
}

/* ===== Responsive ===== */
@media (max-width: 1023px) {
  .page-header { height: 280px; }
  .page-title { font-size: 2.2rem; }

  .story-grid {
    grid-template-columns: 1fr;
    gap: 32px;
  }

  .mission-grid {
    grid-template-columns: 1fr;
    gap: 20px;
  }

  .cert-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 767px) {
  .page-header { height: 160px; padding-top: var(--header-height); }
  .page-title { font-size: 1.5rem; }
  .page-header-content { padding-top: 0; }

  .story-section {
    padding-top: 20px;
    padding-bottom: 24px;
  }

  .mission-section,
  .timeline-section,
  .team-section {
    padding-top: 24px;
    padding-bottom: 24px;
  }

  .story-image img { height: 260px; }

  .timeline::before {
    left: 16px;
  }

  .timeline-item,
  .timeline-item.timeline-right {
    width: 100%;
    margin-left: 0;
    padding-left: 48px;
    padding-right: 0;
  }

  .timeline-dot,
  .timeline-item.timeline-right .timeline-dot {
    left: 8px;
    right: auto;
  }

  .cert-grid {
    grid-template-columns: 1fr;
  }
}
</style>

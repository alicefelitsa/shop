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
              src="https://images.unsplash.com/photo-1532187863486-abf9dbad1b69?w=600&h=500&fit=crop"
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
  height: 360px;
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
  background-image:
    url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 1200 400' preserveAspectRatio='xMidYMid slice'%3E%3Cdefs%3E%3ClinearGradient id='bg' x1='0' y1='0' x2='1' y2='1'%3E%3Cstop offset='0' stop-color='%2316345c'/%3E%3Cstop offset='1' stop-color='%230d2038'/%3E%3C/linearGradient%3E%3CradialGradient id='glow' cx='0.82' cy='0.2' r='0.55'%3E%3Cstop offset='0' stop-color='rgba(255,255,255,0.12)'/%3E%3Cstop offset='1' stop-color='rgba(255,255,255,0)'/%3E%3C/radialGradient%3E%3CradialGradient id='gold' cx='0.12' cy='0.85' r='0.5'%3E%3Cstop offset='0' stop-color='rgba(201,168,76,0.16)'/%3E%3Cstop offset='1' stop-color='rgba(201,168,76,0)'/%3E%3C/radialGradient%3E%3C/defs%3E%3Crect width='1200' height='400' fill='url(%23bg)'/%3E%3Crect width='1200' height='400' fill='url(%23glow)'/%3E%3Crect width='1200' height='400' fill='url(%23gold)'/%3E%3Cg stroke='rgba(255,255,255,0.14)' stroke-width='2' fill='none'%3E%3Cpolygon points='950,90 1002,60 1054,90 1054,150 1002,180 950,150'/%3E%3Cpolygon points='975,104 1002,88 1029,104 1029,136 1002,152 975,136'/%3E%3Cline x1='1054' y1='150' x2='1106' y2='180'/%3E%3Cline x1='950' y1='90' x2='898' y2='60'/%3E%3Cline x1='1002' y1='180' x2='1002' y2='240'/%3E%3Cpolygon points='140,240 192,210 244,240 244,300 192,330 140,300'/%3E%3Cline x1='244' y1='300' x2='296' y2='330'/%3E%3Cline x1='140' y1='240' x2='88' y2='210'/%3E%3Cline x1='400' y1='80' x2='470' y2='50'/%3E%3Cline x1='470' y1='50' x2='540' y2='85'/%3E%3C/g%3E%3Cg%3E%3Ccircle cx='898' cy='60' r='6' fill='rgba(201,168,76,0.6)'/%3E%3Ccircle cx='1106' cy='180' r='7' fill='rgba(255,255,255,0.35)'/%3E%3Ccircle cx='1002' cy='240' r='5' fill='rgba(201,168,76,0.45)'/%3E%3Ccircle cx='88' cy='210' r='5' fill='rgba(255,255,255,0.3)'/%3E%3Ccircle cx='296' cy='330' r='6' fill='rgba(201,168,76,0.5)'/%3E%3Ccircle cx='400' cy='80' r='4' fill='rgba(255,255,255,0.3)'/%3E%3Ccircle cx='470' cy='50' r='5' fill='rgba(201,168,76,0.5)'/%3E%3Ccircle cx='540' cy='85' r='4' fill='rgba(255,255,255,0.25)'/%3E%3Ccircle cx='640' cy='300' r='3' fill='rgba(255,255,255,0.25)'/%3E%3Ccircle cx='760' cy='340' r='3' fill='rgba(201,168,76,0.4)'/%3E%3Ccircle cx='700' cy='120' r='2.5' fill='rgba(255,255,255,0.2)'/%3E%3Ccircle cx='340' cy='150' r='2.5' fill='rgba(255,255,255,0.2)'/%3E%3Ccircle cx='560' cy='200' r='2' fill='rgba(201,168,76,0.35)'/%3E%3Ccircle cx='1140' cy='70' r='3' fill='rgba(255,255,255,0.25)'/%3E%3Ccircle cx='60' cy='350' r='2.5' fill='rgba(255,255,255,0.2)'/%3E%3Ccircle cx='850' cy='300' r='2.5' fill='rgba(255,255,255,0.2)'/%3E%3C/g%3E%3C/svg%3E");
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

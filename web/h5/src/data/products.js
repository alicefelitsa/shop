// Mock product data for the company website
export const categories = [
  { id: 'all', name: 'All Products' },
  { id: 'peptides', name: 'Research Peptides' },
  { id: 'raws', name: 'Peptide Raws' },
  { id: 'equipment', name: 'Lab Equipment' },
  { id: 'blends', name: 'Peptide Blends' }
]

export const products = [
  {
    id: 1,
    name: 'TB-500 (Thymosin Beta-4)',
    slug: 'tb-500',
    category: 'peptides',
    price: 50,
    priceMax: 204,
    originalPrice: null,
    images: [
      'https://images.unsplash.com/photo-1587854692152-cbe660dbde88?w=600&h=600&fit=crop',
      'https://images.unsplash.com/photo-1532187863486-abf9dbad1b69?w=600&h=600&fit=crop'
    ],
    thumbnail: 'https://images.unsplash.com/photo-1587854692152-cbe660dbde88?w=400&h=400&fit=crop',
    description: 'High-purity TB-500 (Thymosin Beta-4) lyophilized peptide for advanced research applications. Known for its role in tissue repair and regeneration studies.',
    specs: { purity: '99.086%', form: 'Lyophilized Powder', storage: '-20°C', molecularWeight: '4963.44 g/mol', sequence: 'Ac-SDKPDMAEIEKFDKSKLKKTETQEKNPLPSKETIEQEKQAGES' },
    variations: [
      { label: '1 Kit', value: '1kit' },
      { label: '5 Kits', value: '5kits' },
      { label: '10 Kits', value: '10kits' }
    ],
    doses: ['2mg/vial', '5mg/vial', '10mg/vial'],
    rating: 4.8,
    reviews: 24,
    featured: true,
    bestSeller: false,
    onSale: false
  },
  {
    id: 2,
    name: 'CJC-1295 (No DAC)',
    slug: 'cjc-1295',
    category: 'peptides',
    price: 47,
    priceMax: 189,
    originalPrice: null,
    images: [
      'https://images.unsplash.com/photo-1576671081837-49000212a370?w=600&h=600&fit=crop',
      'https://images.unsplash.com/photo-1530026405186-ed1f139313f8?w=600&h=600&fit=crop'
    ],
    thumbnail: 'https://images.unsplash.com/photo-1576671081837-49000212a370?w=400&h=400&fit=crop',
    description: 'Premium grade CJC-1295 without DAC for growth hormone releasing research. Widely used in endocrinology and metabolism studies.',
    specs: { purity: '99.24%', form: 'Lyophilized Powder', storage: '-20°C', molecularWeight: '3367.94 g/mol', halfLife: '~30 minutes' },
    variations: [
      { label: '1 Kit', value: '1kit' },
      { label: '5 Kits', value: '5kits' },
      { label: '10 Kits', value: '10kits' }
    ],
    doses: ['2mg/vial', '5mg/vial', '10mg/vial'],
    rating: 4.6,
    reviews: 18,
    featured: true,
    bestSeller: true,
    onSale: false
  },
  {
    id: 3,
    name: 'Tirzepatide Lyophilized Powder',
    slug: 'tirzepatide',
    category: 'peptides',
    price: 57,
    priceMax: 251,
    originalPrice: null,
    images: [
      'https://images.unsplash.com/photo-1585435557343-3b092031a831?w=600&h=600&fit=crop',
      'https://images.unsplash.com/photo-1631549916768-4119b2e5f926?w=600&h=600&fit=crop'
    ],
    thumbnail: 'https://images.unsplash.com/photo-1585435557343-3b092031a831?w=400&h=400&fit=crop',
    description: 'Dual GIP/GLP-1 receptor agonist peptide for metabolic research. Ideal for studies on glucose metabolism, insulin sensitivity, and weight management pathways.',
    specs: { purity: '99.682%', form: 'Lyophilized Powder', storage: '-20°C', molecularWeight: '4813.45 g/mol', mechanism: 'Dual GIP/GLP-1 Agonist' },
    variations: [
      { label: '1 Kit', value: '1kit' },
      { label: '5 Kits', value: '5kits' },
      { label: '10 Kits', value: '10kits' }
    ],
    doses: ['5mg/vial', '10mg/vial', '15mg/vial', '20mg/vial', '30mg/vial'],
    rating: 4.9,
    reviews: 42,
    featured: true,
    bestSeller: true,
    onSale: false
  },
  {
    id: 4,
    name: 'Retatrutide Lyophilized Powder',
    slug: 'retatrutide',
    category: 'peptides',
    price: 71,
    priceMax: 380,
    originalPrice: null,
    images: [
      'https://images.unsplash.com/photo-1614935151651-0bea6508db6b?w=600&h=600&fit=crop',
      'https://images.unsplash.com/photo-1578321272176-b7bbc0679853?w=600&h=600&fit=crop'
    ],
    thumbnail: 'https://images.unsplash.com/photo-1614935151651-0bea6508db6b?w=400&h=400&fit=crop',
    description: 'Triple agonist peptide (GLP-1/GIP/Glucagon receptor) for cutting-edge metabolic research. Next-generation compound for advanced pharmacological studies.',
    specs: { purity: '99.150%', form: 'Lyophilized Powder', storage: '-20°C', molecularWeight: '4167.83 g/mol', mechanism: 'Triple Agonist (GLP-1/GIP/GCG)' },
    variations: [
      { label: '1 Kit', value: '1kit' },
      { label: '5 Kits', value: '5kits' },
      { label: '10 Kits', value: '10kits' }
    ],
    doses: ['5mg/vial', '10mg/vial', '15mg/vial', '20mg/vial', '30mg/vial'],
    rating: 4.7,
    reviews: 31,
    featured: true,
    bestSeller: false,
    onSale: true,
    salePrice: 65
  },
  {
    id: 5,
    name: 'Semaglutide Lyophilized Powder',
    slug: 'semaglutide',
    category: 'peptides',
    price: 41,
    priceMax: 145,
    originalPrice: null,
    images: [
      'https://images.unsplash.com/photo-1579165466991-467135ad3110?w=600&h=600&fit=crop',
      'https://images.unsplash.com/photo-1587854692152-cbe660dbde88?w=600&h=600&fit=crop'
    ],
    thumbnail: 'https://images.unsplash.com/photo-1579165466991-467135ad3110?w=400&h=400&fit=crop',
    description: 'GLP-1 receptor agonist peptide for diabetes and metabolic syndrome research. One of the most studied peptides in current pharmaceutical research.',
    specs: { purity: '99.164%', form: 'Lyophilized Powder', storage: '-20°C', molecularWeight: '4113.64 g/mol', mechanism: 'GLP-1 Receptor Agonist' },
    variations: [
      { label: '1 Kit', value: '1kit' },
      { label: '5 Kits', value: '5kits' }
    ],
    doses: ['2mg/vial', '5mg/vial', '10mg/vial', '15mg/vial'],
    rating: 4.5,
    reviews: 56,
    featured: true,
    bestSeller: true,
    onSale: false
  },
  {
    id: 6,
    name: 'BPC-157 Peptide',
    slug: 'bpc-157',
    category: 'peptides',
    price: 38,
    priceMax: 120,
    originalPrice: null,
    images: [
      'https://images.unsplash.com/photo-1532187863486-abf9dbad1b69?w=600&h=600&fit=crop',
      'https://images.unsplash.com/photo-1587854692152-cbe660dbde88?w=600&h=600&fit=crop'
    ],
    thumbnail: 'https://images.unsplash.com/photo-1532187863486-abf9dbad1b69?w=400&h=400&fit=crop',
    description: 'Body Protection Compound-157 for tissue healing and gastrointestinal research. Extensively studied for its cytoprotective and healing properties.',
    specs: { purity: '99.56%', form: 'Lyophilized Powder', storage: '-20°C', molecularWeight: '1419.53 g/mol', sequence: 'Gly-Glu-Pro-Pro-Pro-Gly-Lys-Pro-Ala-Asp-Asp-Ala-Gly-Leu-Val' },
    variations: [
      { label: '1 Kit', value: '1kit' },
      { label: '5 Kits', value: '5kits' }
    ],
    doses: ['5mg/vial', '10mg/vial'],
    rating: 4.8,
    reviews: 38,
    featured: false,
    bestSeller: true,
    onSale: false
  },
  {
    id: 7,
    name: 'BPC-157 + TB-500 Blend',
    slug: 'bpc-tb-blend',
    category: 'blends',
    price: 141,
    priceMax: 267,
    originalPrice: null,
    images: [
      'https://images.unsplash.com/photo-1550572017-edd951b55104?w=600&h=600&fit=crop',
      'https://images.unsplash.com/photo-1576671081837-49000212a370?w=600&h=600&fit=crop'
    ],
    thumbnail: 'https://images.unsplash.com/photo-1550572017-edd951b55104?w=400&h=400&fit=crop',
    description: 'Research-grade combination blend of BPC-157 and TB-500 for synergistic tissue repair studies. Pre-mixed ratio for convenience in laboratory settings.',
    specs: { purity: '>99%', form: 'Lyophilized Powder', storage: '-20°C', composition: 'BPC-157 10mg + TB-500 10mg per vial', ratio: '1:1' },
    variations: [
      { label: '1 Kit', value: '1kit' },
      { label: '5 Kits', value: '5kits' }
    ],
    doses: ['10mg/vial', '20mg/vial'],
    rating: 4.6,
    reviews: 15,
    featured: true,
    bestSeller: false,
    onSale: false
  },
  {
    id: 8,
    name: 'GHK-Cu Copper Peptide',
    slug: 'ghk-cu',
    category: 'peptides',
    price: 30,
    priceMax: 45,
    originalPrice: null,
    images: [
      'https://images.unsplash.com/photo-1631549916768-4119b2e5f926?w=600&h=600&fit=crop',
      'https://images.unsplash.com/photo-1585435557343-3b092031a831?w=600&h=600&fit=crop'
    ],
    thumbnail: 'https://images.unsplash.com/photo-1631549916768-4119b2e5f926?w=400&h=400&fit=crop',
    description: 'Copper-binding tripeptide for dermatological and wound healing research. Naturally occurring peptide with well-documented regenerative properties.',
    specs: { purity: '>99%', form: 'Lyophilized Powder', storage: '2-8°C', molecularWeight: '403.39 g/mol', copperContent: '~6.2%' },
    variations: [
      { label: '1 Kit', value: '1kit' },
      { label: '5 Kits', value: '5kits' }
    ],
    doses: ['50mg/vial', '100mg/vial'],
    rating: 4.4,
    reviews: 22,
    featured: false,
    bestSeller: false,
    onSale: false
  },
  {
    id: 9,
    name: 'NAD+ (Nicotinamide Adenine Dinucleotide)',
    slug: 'nad-plus',
    category: 'peptides',
    price: 47,
    priceMax: 182,
    originalPrice: null,
    images: [
      'https://images.unsplash.com/photo-1582719471384-894fbb16e074?w=600&h=600&fit=crop',
      'https://images.unsplash.com/photo-1530026405186-ed1f139313f8?w=600&h=600&fit=crop'
    ],
    thumbnail: 'https://images.unsplash.com/photo-1582719471384-894fbb16e074?w=400&h=400&fit=crop',
    description: 'Essential coenzyme for cellular metabolism and aging research. Critical compound in studies of mitochondrial function, DNA repair, and cellular energy.',
    specs: { purity: '>99%', form: 'Lyophilized Powder', storage: '-20°C', molecularWeight: '663.43 g/mol', casNumber: '53-84-9' },
    variations: [
      { label: '1 Kit', value: '1kit' },
      { label: '5 Kits', value: '5kits' }
    ],
    doses: ['100mg/vial', '500mg/vial', '1000mg/vial'],
    rating: 4.7,
    reviews: 29,
    featured: true,
    bestSeller: false,
    onSale: false
  },
  {
    id: 10,
    name: 'IGF-1 LR3',
    slug: 'igf-1-lr3',
    category: 'peptides',
    price: 50,
    priceMax: 245,
    originalPrice: null,
    images: [
      'https://images.unsplash.com/photo-1581093450021-4a7360e9a6b5?w=600&h=600&fit=crop',
      'https://images.unsplash.com/photo-1578321272176-b7bbc0679853?w=600&h=600&fit=crop'
    ],
    thumbnail: 'https://images.unsplash.com/photo-1581093450021-4a7360e9a6b5?w=400&h=400&fit=crop',
    description: 'Long R3 Insulin-like Growth Factor-1 for cell proliferation and growth research. Modified IGF-1 with extended half-life for laboratory studies.',
    specs: { purity: '99.265%', form: 'Lyophilized Powder', storage: '-20°C', molecularWeight: '7649.00 g/mol', halfLife: '12-15 hours' },
    variations: [
      { label: '1 Kit (10 vials)', value: '1kit' }
    ],
    doses: ['0.1mg/vial', '1mg/vial'],
    rating: 4.5,
    reviews: 12,
    featured: false,
    bestSeller: false,
    onSale: false
  },
  {
    id: 11,
    name: 'Semaglutide Raw Powder 1g',
    slug: 'semaglutide-raw',
    category: 'raws',
    price: 2265,
    priceMax: 2369,
    originalPrice: 2369,
    images: [
      'https://images.unsplash.com/photo-1550572017-edd951b55104?w=600&h=600&fit=crop',
      'https://images.unsplash.com/photo-1579165466991-467135ad3110?w=600&h=600&fit=crop'
    ],
    thumbnail: 'https://images.unsplash.com/photo-1550572017-edd951b55104?w=400&h=400&fit=crop',
    description: 'Bulk semaglutide raw powder for pharmaceutical research and development. High-purity grade suitable for analytical and formulation studies.',
    specs: { purity: '99.36%', form: 'Fine Powder', storage: '-20°C, desiccated', packaging: '1g sealed vial', shelfLife: '24 months' },
    variations: [
      { label: '1g', value: '1g' },
      { label: '5g', value: '5g' }
    ],
    doses: [],
    rating: 4.9,
    reviews: 8,
    featured: false,
    bestSeller: true,
    onSale: true,
    salePrice: 2265
  },
  {
    id: 12,
    name: 'Mini Freeze Dryer LY-10',
    slug: 'freeze-dryer-ly10',
    category: 'equipment',
    price: 1568,
    priceMax: 1980,
    originalPrice: null,
    images: [
      'https://images.unsplash.com/photo-1576671081837-49000212a370?w=600&h=600&fit=crop',
      'https://images.unsplash.com/photo-1585435557343-3b092031a831?w=600&h=600&fit=crop'
    ],
    thumbnail: 'https://images.unsplash.com/photo-1576671081837-49000212a370?w=400&h=400&fit=crop',
    description: 'Laboratory-grade mini freeze dryer for peptide and pharmaceutical lyophilization. Compact design with precise temperature control for small-scale production.',
    specs: { capacity: '1-2kg per batch', temperature: '-50°C to +65°C', vacuum: '<10 Pa', power: '220V/50Hz', dimensions: '650×450×720mm' },
    variations: [],
    doses: [],
    rating: 4.3,
    reviews: 6,
    featured: false,
    bestSeller: false,
    onSale: false
  },
  {
    id: 13,
    name: 'Ipamorelin Peptide',
    slug: 'ipamorelin',
    category: 'peptides',
    price: 47,
    priceMax: 82,
    originalPrice: null,
    images: [
      'https://images.unsplash.com/photo-1576671081837-49000212a370?w=600&h=600&fit=crop',
      'https://images.unsplash.com/photo-1587854692152-cbe660dbde88?w=600&h=600&fit=crop'
    ],
    thumbnail: 'https://images.unsplash.com/photo-1576671081837-49000212a370?w=400&h=400&fit=crop',
    description: 'Selective growth hormone secretagogue for endocrinology research. Pentapeptide with well-characterized GH-releasing properties and minimal side effects in studies.',
    specs: { purity: '99.61%', form: 'Lyophilized Powder', storage: '-20°C', molecularWeight: '711.85 g/mol', sequence: 'Aib-His-D-2-Nal-D-Phe-Lys-NH2' },
    variations: [
      { label: '1 Kit', value: '1kit' },
      { label: '5 Kits', value: '5kits' }
    ],
    doses: ['5mg/vial', '10mg/vial'],
    rating: 4.6,
    reviews: 19,
    featured: false,
    bestSeller: false,
    onSale: false
  },
  {
    id: 14,
    name: 'AOD-9604 Peptide',
    slug: 'aod-9604',
    category: 'peptides',
    price: 63,
    priceMax: 236,
    originalPrice: null,
    images: [
      'https://images.unsplash.com/photo-1582719471384-894fbb16e074?w=600&h=600&fit=crop',
      'https://images.unsplash.com/photo-1614935151651-0bea6508db6b?w=600&h=600&fit=crop'
    ],
    thumbnail: 'https://images.unsplash.com/photo-1582719471384-894fbb16e074?w=400&h=400&fit=crop',
    description: 'Modified fragment of human growth hormone for fat metabolism research. Specifically designed for studying lipid metabolism and adipocyte regulation.',
    specs: { purity: '>99%', form: 'Lyophilized Powder', storage: '-20°C', molecularWeight: '1785.11 g/mol', fragment: 'hGH 177-191 modified' },
    variations: [
      { label: '1 Kit', value: '1kit' },
      { label: '5 Kits', value: '5kits' }
    ],
    doses: ['2mg/vial', '5mg/vial', '10mg/vial'],
    rating: 4.4,
    reviews: 14,
    featured: false,
    bestSeller: false,
    onSale: false
  },
  {
    id: 15,
    name: 'CellRestore Research Blend',
    slug: 'cellrestore-blend',
    category: 'blends',
    price: 283,
    priceMax: 388,
    originalPrice: 388,
    images: [
      'https://images.unsplash.com/photo-1631549916768-4119b2e5f926?w=600&h=600&fit=crop',
      'https://images.unsplash.com/photo-1550572017-edd951b55104?w=600&h=600&fit=crop'
    ],
    thumbnail: 'https://images.unsplash.com/photo-1631549916768-4119b2e5f926?w=400&h=400&fit=crop',
    description: 'Advanced triple-peptide complex combining BPC-157, TB-500, and GHK-Cu for comprehensive tissue regeneration research. Premium research formulation.',
    specs: { purity: '>99% each', form: 'Lyophilized Powder', storage: '-20°C', composition: 'BPC-157 + TB-500 + GHK-Cu', type: 'Triple Blend' },
    variations: [],
    doses: [],
    rating: 4.8,
    reviews: 9,
    featured: true,
    bestSeller: false,
    onSale: true,
    salePrice: 283
  },
  {
    id: 16,
    name: 'Bacteriostatic Water (BAC Water)',
    slug: 'bac-water',
    category: 'equipment',
    price: 10,
    priceMax: 15,
    originalPrice: null,
    images: [
      'https://images.unsplash.com/photo-1579165466991-467135ad3110?w=600&h=600&fit=crop',
      'https://images.unsplash.com/photo-1530026405186-ed1f139313f8?w=600&h=600&fit=crop'
    ],
    thumbnail: 'https://images.unsplash.com/photo-1579165466991-467135ad3110?w=400&h=400&fit=crop',
    description: 'Sterile bacteriostatic water for peptide reconstitution and research use. USP-grade with 0.9% benzyl alcohol preservative.',
    specs: { grade: 'USP', form: 'Sterile Solution', storage: 'Room Temperature', preservative: '0.9% Benzyl Alcohol', shelfLife: '28 days after opening' },
    variations: [
      { label: '1 Vial', value: '1vial' },
      { label: '10 Vials', value: '10vials' }
    ],
    doses: ['3ml/vial', '10ml/vial'],
    rating: 4.2,
    reviews: 45,
    featured: false,
    bestSeller: false,
    onSale: false
  }
]

export const banners = [
  {
    id: 1,
    title: 'Premium Research Peptides',
    subtitle: 'Trusted by researchers worldwide. Lab-tested purity above 99%.',
    image: 'https://images.unsplash.com/photo-1532187863486-abf9dbad1b69?w=1920&h=900&fit=crop',
    cta: 'Explore Products',
    link: '/products'
  },
  {
    id: 2,
    title: 'Cutting-Edge Lab Equipment',
    subtitle: 'Professional-grade machinery for peptide synthesis and lyophilization.',
    image: 'https://images.unsplash.com/photo-1585435557343-3b092031a831?w=1920&h=900&fit=crop',
    cta: 'View Equipment',
    link: '/products?category=equipment'
  },
  {
    id: 3,
    title: 'Wholesale & Bulk Orders',
    subtitle: 'Competitive pricing for research institutions and laboratories.',
    image: 'https://images.unsplash.com/photo-1587854692152-cbe660dbde88?w=1920&h=900&fit=crop',
    cta: 'Contact Us',
    link: '/contact'
  }
]

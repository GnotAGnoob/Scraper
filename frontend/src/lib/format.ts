export const formatPrice = (value?: number, currency = 'CZK') => {
  if (typeof value !== 'number' || Number.isNaN(value)) {
    return '—'
  }

  return new Intl.NumberFormat('cs-CZ', {
    style: 'currency',
    currency,
    maximumFractionDigits: 2,
  }).format(value)
}

export const formatQuantity = (grams?: number) => {
  if (!grams || Number.isNaN(grams)) {
    return 'as needed'
  }

  if (grams < 1000) {
    return `${Math.round(grams)} g`
  }

  return `${(grams / 1000).toFixed(1)} kg`
}



import FontType, { declaredFontTypes } from 'src/FontType';

test('descriptive names', async () => {
  expect(declaredFontTypes('')).toEqual([]);
  expect(declaredFontTypes('JBRMKS+Helvetica')).toEqual([]);
  expect(declaredFontTypes('OMUGKQ+Helvetica-Bold')).toEqual([FontType.BOLD]);
  expect(declaredFontTypes('SVUOCV+Helvetica-Oblique')).toEqual([FontType.OBLIQUE]);
  expect(declaredFontTypes('JUJONH+Helvetica-BoldOblique')).toEqual([FontType.BOLD, FontType.OBLIQUE]);
});

// See http://mirrors.ibiblio.org/CTAN/systems/win32/bakoma/fonts/fonts.html
test('ATM Compatible Postscript Type 1', async () => {
  expect(declaredFontTypes('')).toEqual([]);
  expect(declaredFontTypes('BBXMCN+CMR9')).toEqual([]);
  expect(declaredFontTypes('EFUEQI+CMR10')).toEqual([]);
  expect(declaredFontTypes('JZXNAL+CMCSC10')).toEqual([]);
  expect(declaredFontTypes('ZYSMDY+CMBX10')).toEqual([FontType.BOLD]);
  expect(declaredFontTypes('AENRCE+CMBX12')).toEqual([FontType.BOLD]);
  expect(declaredFontTypes('HENPPA+BitstreamCyberbit-Roman')).toEqual([]);
  expect(declaredFontTypes('GHPDYG+CMSY10')).toEqual([]);
  expect(declaredFontTypes('VKLUIG+CMTT9')).toEqual([]);
  expect(declaredFontTypes('KSVJZ+CMTI10')).toEqual([FontType.OBLIQUE]);
  expect(declaredFontTypes('QCQOVJ+CMTT10')).toEqual([]);
  expect(declaredFontTypes('ASZLVZ+BitstreamCyberbit-Roman')).toEqual([]);
  expect(declaredFontTypes('KFYFQJ+CMMI10')).toEqual([FontType.OBLIQUE]);
  expect(declaredFontTypes('GYUWCJ+CMMIB10')).toEqual([FontType.BOLD, FontType.OBLIQUE]);
  expect(declaredFontTypes('OUVHFK+CMR8')).toEqual([]);
});

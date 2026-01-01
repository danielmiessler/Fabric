import { test, expect } from '@playwright/test';
import { describe } from 'node:test';

describe('default settings', () => {
  test('correctly changes the element when it enters the viewport', async ({
    page,
  }) => {
    await page.goto('/default-settings');
    const targetBlock = page.locator('.target-block');

    await expect(targetBlock).toBeDefined();
    await expect(targetBlock).toHaveText(/no/i);

    await page.evaluate(() => window.scroll(0, 150));

    await expect(targetBlock).toHaveText(/yes/i);
  });

  test('correctly changes back the element when it leaves the viewport', async ({
    page,
  }) => {
    await page.goto('/default-settings');
    const targetBlock = page.locator('.target-block');

    await expect(targetBlock).toBeDefined();
    await expect(targetBlock).toHaveText(/no/i);

    await page.evaluate(() => window.scroll(0, 150));

    await expect(targetBlock).toHaveText(/yes/i);

    await page.evaluate(() => window.scroll(0, 1500));

    await expect(targetBlock).toHaveText(/no/i);
  });
});

describe('custom settings', () => {
  test('correctly changes the element with custom rootMargin setting', async ({
    page,
  }) => {
    await page.goto('/root-margin');
    const targetBlock = page.locator('.target-block');

    await expect(targetBlock).toBeDefined();
    await expect(targetBlock).toHaveText(/no/i);

    await page.evaluate(() => window.scroll(0, 150));

    await expect(targetBlock).toHaveText(/no/i);

    await page.evaluate(() => window.scroll(0, 250));

    await expect(targetBlock).toHaveText(/yes/i);
  });

  test('correctly changes the element with custom threshold setting', async ({
    page,
  }) => {
    await page.goto('/threshold');
    const targetBlock = page.locator('.target-block');

    await expect(targetBlock).toBeDefined();
    await expect(targetBlock).toHaveText(/no/i);

    await page.evaluate(() => window.scroll(0, 150));

    await expect(targetBlock).toHaveText(/no/i);

    await page.evaluate(() => window.scroll(0, 250));

    await expect(targetBlock).toHaveText(/yes/i);
  });

  test('does not change the element back when it leaves the viewport with custom unobserveOnEnter setting', async ({
    page,
  }) => {
    await page.goto('/unobserve-on-enter');

    const targetBlock = page.locator('.target-block');

    await expect(targetBlock).toBeDefined();
    await expect(targetBlock).toHaveText(/no/i);

    await page.evaluate(() => window.scroll(0, 250));

    await expect(targetBlock).toHaveText(/yes/i);

    await page.evaluate(() => window.scroll(0, 1500));

    await expect(targetBlock).toHaveText(/yes/i);
  });
});

describe('direction', () => {
  test('shows the correct direction', async ({ page }) => {
    await page.goto('/direction');
    const targetBlock = page.locator('.target-block');

    await expect(targetBlock).toBeDefined();
    await expect(targetBlock).toHaveText(/down/i);

    await page.evaluate(() => window.scroll(0, 250));

    await expect(targetBlock).toHaveText(/up/i);

    await page.evaluate(() => window.scroll(0, 1500));
    await page.evaluate(() => window.scroll(0, -1));

    await expect(targetBlock).toHaveText(/no/i);
  });
});

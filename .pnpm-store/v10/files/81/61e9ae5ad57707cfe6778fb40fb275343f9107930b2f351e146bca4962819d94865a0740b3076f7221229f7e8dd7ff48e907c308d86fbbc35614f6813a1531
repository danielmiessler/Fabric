import Progress from 'src/Progress';

test('basic progress', async () => {
  const progress = new Progress(['Stage0', 'Stage1', 'Stage1']);

  // nothing yet
  expectTotalProgress(progress, 0);
  expectStageInProgress(progress, 0);

  // stage 0 progress
  progress.stageProgress[0] = 0.3;
  expectTotalProgress(progress, 10);
  expectStageInProgress(progress, 0);

  // stage 0 completed
  progress.stageProgress[0] = 1;
  expectTotalProgress(progress, 33);
  expectStageInProgress(progress, 1);

  // stage 1 progress
  progress.stageProgress[1] = 0.3;
  expectTotalProgress(progress, 43);
  expectStageInProgress(progress, 1);

  // stage 1 completed
  progress.stageProgress[1] = 1;
  expectTotalProgress(progress, 67);
  expectStageInProgress(progress, 2);

  // stage 2 completed
  progress.stageProgress[2] = 1;
  expectTotalProgress(progress, 100);
  expectStageInProgress(progress, 3);
});

test('number of stage weights must match the number of stages', async () => {
  try {
    new Progress(['Stage0', 'Stage1', 'Stage1'], [0.5, 0.5]);
    fail('Creating a progress object with number of weigths not matching numbers of stages should fail');
  } catch (error) {
    expect(error.message).toEqual('Provided only 2 weights but expected 3 for 3 stages');
  }
});

test('stage weights must sum up', async () => {
  try {
    new Progress(['Stage0', 'Stage1', 'Stage1'], [0.5, 0.5, 0.5]);
    fail('Creating a progress object with stage weigths not summing up should fail');
  } catch (error) {
    expect(error.message).toEqual('Weights [0.5, 0.5, 0.5] should sum up to 1, but did to 1.5');
  }
});

test('weighted progress', async () => {
  const progress = new Progress(['Stage0', 'Stage1', 'Stage1'], [0, 0.7, 0.3]);

  // nothing yet
  expectTotalProgress(progress, 0);

  // stage 0 progress
  progress.stageProgress[0] = 0.9;
  expectTotalProgress(progress, 0);

  // stage 0 completed
  progress.stageProgress[0] = 1;
  expectTotalProgress(progress, 0);

  // stage 1 progress
  progress.stageProgress[1] = 0.3;
  expectTotalProgress(progress, 21);

  // stage 1 more progress
  progress.stageProgress[1] = 0.6;
  expectTotalProgress(progress, 42);

  // stage 1 completed
  progress.stageProgress[1] = 1;
  expectTotalProgress(progress, 70);

  // stage 2 progress
  progress.stageProgress[2] = 0.3;
  expectTotalProgress(progress, 79);

  // stage 2 completed
  progress.stageProgress[2] = 1;
  expectTotalProgress(progress, 100);
});

function expectTotalProgress(progress: Progress, expected: number) {
  expect(Math.round(progress.totalProgress() * 100)).toBe(expected);
}

function expectStageInProgress(progress: Progress, stageIndex: number) {
  for (let index = 0; index < progress.stageProgress.length; index++) {
    if (index < stageIndex) {
      expect(progress.isProgressing(index)).toBe(false);
      expect(progress.isComplete(index)).toBe(true);
    } else if (index === stageIndex) {
      expect(progress.isProgressing(index)).toBe(true);
      expect(progress.isComplete(index)).toBe(false);
    } else if (index > stageIndex) {
      expect(progress.isProgressing(index)).toBe(false);
      expect(progress.isComplete(index)).toBe(false);
    }
  }
}

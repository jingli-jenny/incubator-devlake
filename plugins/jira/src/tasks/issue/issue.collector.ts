import { Inject } from '@nestjs/common';
import Exports from 'plugins/core/src/exports.decorator';
import Task from 'plugins/core/src/task.interface';
import { Repository } from 'typeorm';
import IssueEntity from './issue.entity';

export type JiraSource = {
  host: string;
  username: string;
  token: string;
};

@Exports(IssueEntity)
export default class IssueCollector implements Task {
  @Inject(IssueEntity) private IssueRepository: Repository<IssueEntity>;

  name(): string {
    return 'JiraIssue';
  }

  async execute(source: JiraSource): Promise<void> {
    //TODO: do collector
    console.info('Excute Jira Issue Collector', source);
    return;
  }
}
// internal/dynamodb/task_repository.go
package dynamodb

import (
    "context"
    "encoding/json"
    "go-graphql-taskflow/internal/model"
    "go-graphql-taskflow/internal/repository"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoTaskRepository struct {
    Client *dynamodb.Client
    Table  string
}

func NewDynamoTaskRepository(client *dynamodb.Client, table string) repository.TaskRepository {
    return &DynamoTaskRepository{
        Client: client,
        Table:  table,
    }
}

func (r *DynamoTaskRepository) CreateTask(ctx context.Context, task *model.Task) error {
    item, _ := json.Marshal(task)
    _, err := r.Client.PutItem(ctx, &dynamodb.PutItemInput{
        TableName: &r.Table,
        Item: map[string]types.AttributeValue{
            "PK":         &types.AttributeValueMemberS{Value: "PROJECT#" + task.ProjectID},
            "SK":         &types.AttributeValueMemberS{Value: "TASK#" + task.ID},
            "EntityData": &types.AttributeValueMemberS{Value: string(item)},
        },
    })
    return err
}

func (r *DynamoTaskRepository) GetTasks(ctx context.Context, projectID string) ([]*model.Task, error) {
    resp, err := r.Client.Query(ctx, &dynamodb.QueryInput{
        TableName:              &r.Table,
        KeyConditionExpression: aws.String("PK = :pk"),
        ExpressionAttributeValues: map[string]types.AttributeValue{
            ":pk": &types.AttributeValueMemberS{Value: "PROJECT#" + projectID},
        },
    })
    if err != nil {
        return nil, err
    }
    tasks := []*model.Task{}
    for _, item := range resp.Items {
        var t model.Task
        json.Unmarshal([]byte(item["EntityData"].(*types.AttributeValueMemberS).Value), &t)
        tasks = append(tasks, &t)
    }
    return tasks, nil
}
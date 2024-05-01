# Spec

```typescript
interface User {
  id: UUID;
  username: string;
  encrypted_password: string;
  email: string;
  is_admin: boolean; // admin users can manage exercises
}

interface Exercise {
  id: UUID;
  name: string;
  uri: string;
  description: string;
  image: Image;
  thumbnail_image: Image;

  // verified exercises show up in search
  verified: boolean;
  verified_by: User;
  verified_at: DateTime;
}

// TODO: Add feature trackable training metrics
// interface TrainingMetric {
//    exercise: Exercise;
//    user: User;
//    set_history: Set[]
// }

// TODO: Add feature user training settings
// interface UserTrianingSettings {
//   granularity: 0.125 | 0.25 | 0.5 | 1 | 1.25; // smallest weight increment allowed
//   unit_type: "kg" | "lbs";
//   user: User;
// }

interface WorkoutSessionTemplate {
  id: UUID;
  name: string;
  description: string;
  uri: string;
  set_group_templates: SetGroupTemplate[];
}

interface WorkoutSession {
  start_at: Date;
  end_at: Date;
  user: User;
  workout_session_template: WorkoutSessionTemplate;
  user_set_groups: UserSetGroup[];
}

interface SetGroupTemplate {
  rest_after_working_sets_in_seconds: number;
  set_group_type: "straight" | "superset";

  set_templates: SetTemplate[];
}

interface SetGroup {
  set_group_template: SetGroupTemplate;
  sets: Set[];
}

interface SetTemplate {
  id: UUID;
  position: number;
  set_type: "warm_up" | "working_set" | "backoff_set";
  percentage_of_one_rep_max: number;
  rate_of_perceived_exertion: number; // rate of perceived exertion [1-10]
  reps_in_reserve: number; // [0-5]
  weight: number;
  reps: number;
  note: string;

  //relations
  exercise: Exercise;
}

interface Set extends SetPlan {
  id: UUID;
  set_result: "success" | "success_but_form_breakdown" | "failure";
  weight: number;
  reps: number;
  rate_of_perceived_exertion: number; // rate of perceived exertion [1-10]
  reps_in_reserve: number; // [0-5]
  note: string;

  // relations
  user: User;
  exercise: Exercise;
  set_template: SetTemplate; // how do we handle deleted set plans
}
```

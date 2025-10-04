create extension if not exists "pgcrypto";

-- Sessions store conversational metadata
create table if not exists public.sessions (
    id uuid primary key default gen_random_uuid(),
    user_id uuid,
    title text not null,
    description text,
    metadata jsonb default '{}'::jsonb,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

-- Messages are linked to sessions and ordered chronologically
create table if not exists public.messages (
    id uuid primary key default gen_random_uuid(),
    session_id uuid not null references public.sessions(id) on delete cascade,
    role text not null check (role in ('system', 'user', 'assistant', 'tool')),
    content text not null,
    metadata jsonb default '{}'::jsonb,
    created_at timestamptz not null default now()
);

-- Patterns persist reusable prompt templates
create table if not exists public.patterns (
    id uuid primary key default gen_random_uuid(),
    name text not null unique,
    description text,
    body text not null,
    tags text[] default array[]::text[],
    is_system boolean not null default false,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

-- Notes capture user annotations or learnings from a session
create table if not exists public.notes (
    id uuid primary key default gen_random_uuid(),
    session_id uuid references public.sessions(id) on delete set null,
    title text not null,
    content text not null,
    tags text[] default array[]::text[],
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

-- Optional lightweight user profile table for personalization
create table if not exists public.users (
    id uuid primary key,
    email text unique,
    display_name text,
    avatar_url text,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

create index if not exists idx_messages_session_id_created_at
    on public.messages(session_id, created_at);

create index if not exists idx_notes_session_id
    on public.notes(session_id);

create index if not exists idx_patterns_tags
    on public.patterns using gin(tags);

alter table public.sessions
    add constraint sessions_user_fk foreign key (user_id) references public.users(id) on delete set null;

-- Trigger to keep updated_at current
create or replace function public.set_updated_at()
returns trigger as $$
begin
  new.updated_at = now();
  return new;
end;
$$ language plpgsql;

create trigger sessions_updated_at before update on public.sessions
    for each row execute function public.set_updated_at();

create trigger patterns_updated_at before update on public.patterns
    for each row execute function public.set_updated_at();

create trigger notes_updated_at before update on public.notes
    for each row execute function public.set_updated_at();

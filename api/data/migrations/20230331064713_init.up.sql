CREATE TABLE IF NOT EXISTS public.categories
(
    id          BIGINT PRIMARY KEY,
    external_id TEXT                     NOT NULL CONSTRAINT categories_external_id_check CHECK (external_id <> ''::text),
    name        TEXT                     NOT NULL CONSTRAINT categories_name_check CHECK (name <> ''::text),
    description TEXT                     NOT NULL CONSTRAINT categories_description_check CHECK (description <> ''::text),
    status      TEXT                     NOT NULL CONSTRAINT categories_status_check CHECK (status <> ''::text),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
    );
CREATE UNIQUE INDEX IF NOT EXISTS categories_uidx_external_id ON public.categories (external_id);
CREATE        INDEX IF NOT EXISTS categories_uidx_name ON public.categories (name);

CREATE TABLE IF NOT EXISTS public.products
(
    id          BIGINT PRIMARY KEY,
    external_id TEXT                     NOT NULL CONSTRAINT products_external_id_check CHECK (external_id <> ''::text),
    name        TEXT                     NOT NULL CONSTRAINT products_name_check CHECK (name <> ''::text),
    description TEXT                     NOT NULL CONSTRAINT products_description_check CHECK (description <> ''::text),
    status      TEXT                     NOT NULL CONSTRAINT products_status_check CHECK (status <> ''::text),
    price       BIGINT                   NOT NULL CONSTRAINT products_price_check CHECK (price > 0),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
    );
CREATE UNIQUE INDEX IF NOT EXISTS products_uidx_external_id ON public.products (external_id);
CREATE        INDEX IF NOT EXISTS products_uidx_name ON public.products (name);

CREATE TABLE IF NOT EXISTS public.product_categories
(
    id          BIGINT PRIMARY KEY,
    product_id  BIGINT                   NOT NULL CONSTRAINT product_categories_product_id_fkey REFERENCES products (id),
    category_id BIGINT                   NOT NULL CONSTRAINT product_categories_category_id_fky REFERENCES categories (id),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
    );
